package controllers

import (
	"context"

	kubeegressgatewayv1alpha1 "github.com/Azure/kube-egress-gateway/api/v1alpha1"
	"github.com/Azure/kube-egress-gateway/pkg/utils/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	testName      = "test"
	testNamespace = "testns"
	privK         = "GHuMwljFfqd2a7cs6BaUOmHflK23zME8VNvC5B37S3k="
	pubK          = "aPxGwq8zERHQ3Q1cOZFdJ+cvJX5Ka4mLN38AyYKYF10="
)

var _ = Describe("StaticGatewayConfiguration controller unit tests", func() {
	var (
		s = scheme.Scheme
		r *StaticGatewayConfigurationReconciler
	)

	Context("Reconcile", func() {
		var (
			req           reconcile.Request
			res           reconcile.Result
			cl            client.Client
			reconcileErr  error
			getErr        error
			gwConfig      *kubeegressgatewayv1alpha1.StaticGatewayConfiguration
			foundGWConfig = &kubeegressgatewayv1alpha1.StaticGatewayConfiguration{}
		)

		BeforeEach(func() {
			req = reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      testName,
					Namespace: testNamespace,
				},
			}
			gwConfig = &kubeegressgatewayv1alpha1.StaticGatewayConfiguration{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testName,
					Namespace: testNamespace,
				},
				Spec: kubeegressgatewayv1alpha1.StaticGatewayConfigurationSpec{
					GatewayNodepoolName: "testgw",
					GatewayVMSSProfile: kubeegressgatewayv1alpha1.GatewayVMSSProfile{
						VMSSResourceGroup:  "vmssRG",
						VMSSName:           "vmss",
						PublicIPPrefixSize: 31,
					},
					PublicIpPrefixId: "testPipPrefix",
				},
			}
			s.AddKnownTypes(kubeegressgatewayv1alpha1.GroupVersion, gwConfig,
				&kubeegressgatewayv1alpha1.GatewayLBConfiguration{},
				&kubeegressgatewayv1alpha1.GatewayVMConfiguration{})
		})

		When("gwConfig is not found", func() {
			BeforeEach(func() {
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("should only report error in get", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(apierrors.IsNotFound(getErr)).To(BeTrue())
				Expect(res).To(Equal(ctrl.Result{}))
			})
		})

		When("gwConfig is newly created", func() {
			BeforeEach(func() {
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should add finalizer", func() {
				Expect(controllerutil.ContainsFinalizer(foundGWConfig, SGCFinalizerName)).To(BeTrue())
			})
		})

		When("gwConfig only has finalizer", func() {
			BeforeEach(func() {
				controllerutil.AddFinalizer(gwConfig, SGCFinalizerName)
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should create a new secret", func() {
				secret := &corev1.Secret{}
				err := getResource(cl, secret)
				Expect(err).To(BeNil())

				privateKeyBytes, ok := secret.Data[WireguardSecretKeyName]
				Expect(ok).To(BeTrue())
				Expect(privateKeyBytes).NotTo(BeEmpty())

				existing := metav1.GetControllerOf(secret)
				Expect(existing).NotTo(BeNil())
				Expect(existing.Name).To(Equal(testName))

				Expect(foundGWConfig.Status.WireguardPrivateKeySecretRef).NotTo(BeNil())
				Expect(foundGWConfig.Status.WireguardPrivateKeySecretRef.Name).To(Equal(testName))

				wgPrivateKey, err := wgtypes.ParseKey(string(privateKeyBytes))
				Expect(err).To(BeNil())
				wgPublicKey := wgPrivateKey.PublicKey().String()
				Expect(foundGWConfig.Status.WireguardPublicKey).To(Equal(wgPublicKey))
			})

			It("should create a new lbConfig", func() {
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{}
				err := getResource(cl, lbConfig)
				Expect(err).To(BeNil())

				Expect(lbConfig.Spec.GatewayNodepoolName).To(Equal(gwConfig.Spec.GatewayNodepoolName))
				Expect(lbConfig.Spec.GatewayVMSSProfile).To(Equal(gwConfig.Spec.GatewayVMSSProfile))

				existing := metav1.GetControllerOf(lbConfig)
				Expect(existing).NotTo(BeNil())
				Expect(existing.Name).To(Equal(testName))

				Expect(foundGWConfig.Status.WireguardServerIP).To(BeEmpty())
				Expect(foundGWConfig.Status.WireguardServerPort).To(BeZero())
			})

			It("should create a new vmConfig", func() {
				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{}
				err := getResource(cl, vmConfig)
				Expect(err).To(BeNil())

				Expect(vmConfig.Spec.GatewayNodepoolName).To(Equal(gwConfig.Spec.GatewayNodepoolName))
				Expect(vmConfig.Spec.GatewayVMSSProfile).To(Equal(gwConfig.Spec.GatewayVMSSProfile))
				Expect(vmConfig.Spec.PublicIpPrefixId).To(Equal(vmConfig.Spec.PublicIpPrefixId))

				existing := metav1.GetControllerOf(vmConfig)
				Expect(existing).NotTo(BeNil())
				Expect(existing.Name).To(Equal(testName))

				Expect(foundGWConfig.Status.PublicIpPrefix).To(BeEmpty())
			})
		})

		When("secret, lbConfig, vmConfig can all be found with status", func() {
			BeforeEach(func() {
				controllerutil.AddFinalizer(gwConfig, SGCFinalizerName)
				secret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Data: map[string][]byte{
						WireguardSecretKeyName: []byte(privK),
					},
				}
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Spec: kubeegressgatewayv1alpha1.GatewayLBConfigurationSpec{
						GatewayNodepoolName: "testgw",
						GatewayVMSSProfile: kubeegressgatewayv1alpha1.GatewayVMSSProfile{
							VMSSResourceGroup:  "vmssRG",
							VMSSName:           "vmss",
							PublicIPPrefixSize: 31,
						},
					},
					Status: kubeegressgatewayv1alpha1.GatewayLBConfigurationStatus{
						FrontendIP: "1.1.1.1",
						ServerPort: 6000,
					},
				}
				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Spec: kubeegressgatewayv1alpha1.GatewayVMConfigurationSpec{
						GatewayNodepoolName: "testgw",
						GatewayVMSSProfile: kubeegressgatewayv1alpha1.GatewayVMSSProfile{
							VMSSResourceGroup:  "vmssRG",
							VMSSName:           "vmss",
							PublicIPPrefixSize: 31,
						},
						PublicIpPrefixId: "testPipPrefix",
					},
					Status: kubeegressgatewayv1alpha1.GatewayVMConfigurationStatus{
						EgressIpPrefix: "1.2.3.4/31",
					},
				}
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig, secret, lbConfig, vmConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should update gwConfig's status from secret, lbConfig and vmConfig", func() {
				Expect(foundGWConfig.Status.WireguardPublicKey).To(Equal(pubK))
				Expect(foundGWConfig.Status.WireguardServerIP).To(Equal("1.1.1.1"))
				Expect(foundGWConfig.Status.WireguardServerPort).To(Equal(int32(6000)))
				Expect(foundGWConfig.Status.PublicIpPrefix).To(Equal("1.2.3.4/31"))
			})
		})

		When("updating gwConfig", func() {
			BeforeEach(func() {
				controllerutil.AddFinalizer(gwConfig, SGCFinalizerName)
				secret := &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Data: map[string][]byte{
						WireguardSecretKeyName: []byte(privK),
					},
				}
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Spec: kubeegressgatewayv1alpha1.GatewayLBConfigurationSpec{
						GatewayNodepoolName: "testgw1",
						GatewayVMSSProfile: kubeegressgatewayv1alpha1.GatewayVMSSProfile{
							VMSSResourceGroup:  "vmssRG1",
							VMSSName:           "vmss1",
							PublicIPPrefixSize: 30,
						},
					},
					Status: kubeegressgatewayv1alpha1.GatewayLBConfigurationStatus{
						FrontendIP: "1.1.1.1",
						ServerPort: 6000,
					},
				}
				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
					Spec: kubeegressgatewayv1alpha1.GatewayVMConfigurationSpec{
						GatewayNodepoolName: "testgw1",
						GatewayVMSSProfile: kubeegressgatewayv1alpha1.GatewayVMSSProfile{
							VMSSResourceGroup:  "vmssRG1",
							VMSSName:           "vmss1",
							PublicIPPrefixSize: 30,
						},
						PublicIpPrefixId: "testPipPrefix1",
					},
					Status: kubeegressgatewayv1alpha1.GatewayVMConfigurationStatus{
						EgressIpPrefix: "1.2.3.4/31",
					},
				}
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig, secret, lbConfig, vmConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should update lbConfig and vmConfig accordingly", func() {
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{}
				err := getResource(cl, lbConfig)
				Expect(err).To(BeNil())

				Expect(lbConfig.Spec.GatewayNodepoolName).To(Equal(gwConfig.Spec.GatewayNodepoolName))
				Expect(lbConfig.Spec.GatewayVMSSProfile).To(Equal(gwConfig.Spec.GatewayVMSSProfile))

				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{}
				err = getResource(cl, vmConfig)
				Expect(err).To(BeNil())

				Expect(vmConfig.Spec.GatewayNodepoolName).To(Equal(gwConfig.Spec.GatewayNodepoolName))
				Expect(vmConfig.Spec.GatewayVMSSProfile).To(Equal(gwConfig.Spec.GatewayVMSSProfile))
				Expect(vmConfig.Spec.PublicIpPrefixId).To(Equal(vmConfig.Spec.PublicIpPrefixId))
			})
		})

		When("deleting a gwConfig without finalizer", func() {
			BeforeEach(func() {
				gwConfig.ObjectMeta.DeletionTimestamp = to.Ptr(metav1.Now())
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error and should do nothing", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(getErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})
		})

		When("deleting a gwConfig with finalizer but no subresources", func() {
			BeforeEach(func() {
				gwConfig.ObjectMeta.DeletionTimestamp = to.Ptr(metav1.Now())
				controllerutil.AddFinalizer(gwConfig, SGCFinalizerName)
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should delete gwConfig", func() {
				Expect(apierrors.IsNotFound(getErr)).To(BeTrue())
			})
		})

		When("deleting a gwConfig with finalizer and subresources", func() {
			BeforeEach(func() {
				gwConfig.ObjectMeta.DeletionTimestamp = to.Ptr(metav1.Now())
				controllerutil.AddFinalizer(gwConfig, SGCFinalizerName)
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
				}
				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{
					ObjectMeta: metav1.ObjectMeta{
						Name:      testName,
						Namespace: testNamespace,
					},
				}
				controllerutil.AddFinalizer(lbConfig, LBConfigFinalizerName)
				controllerutil.AddFinalizer(vmConfig, VMConfigFinalizerName)
				cl = fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects(gwConfig, lbConfig, vmConfig).Build()
				r = &StaticGatewayConfigurationReconciler{Client: cl, Scheme: s}
				res, reconcileErr = r.Reconcile(context.TODO(), req)
				getErr = getResource(cl, foundGWConfig)
			})

			It("shouldn't error", func() {
				Expect(reconcileErr).To(BeNil())
				Expect(res).To(Equal(ctrl.Result{}))
			})

			It("should not delete gwConfig", func() {
				Expect(getErr).To(BeNil())
				Expect(controllerutil.ContainsFinalizer(foundGWConfig, SGCFinalizerName)).To(BeTrue())
			})

			It("should delete subresources", func() {
				lbConfig := &kubeegressgatewayv1alpha1.GatewayLBConfiguration{}
				err := getResource(cl, lbConfig)
				Expect(err).To(BeNil())

				vmConfig := &kubeegressgatewayv1alpha1.GatewayVMConfiguration{}
				err = getResource(cl, vmConfig)
				Expect(err).To(BeNil())

				Expect(lbConfig.GetDeletionTimestamp().IsZero()).To(BeFalse())
				Expect(vmConfig.GetDeletionTimestamp().IsZero()).To(BeFalse())
			})
		})
	})
})

func getResource(cl client.Client, object client.Object) error {
	key := types.NamespacedName{
		Name:      testName,
		Namespace: testNamespace,
	}
	err := cl.Get(context.TODO(), key, object)
	return err
}