// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	bean "github.com/devtron-labs/devtron/api/bean"
	application "github.com/devtron-labs/devtron/client/k8s/application"

	client "github.com/devtron-labs/devtron/api/helm-app"

	cluster "github.com/devtron-labs/devtron/pkg/cluster"

	context "context"

	http "net/http"

	io "io"

	k8s "github.com/devtron-labs/devtron/util/k8s"

	kubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	terminal "github.com/devtron-labs/devtron/pkg/terminal"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	util "github.com/devtron-labs/devtron/internal/util"

	version "k8s.io/apimachinery/pkg/version"
)

// K8sApplicationService is an autogenerated mock type for the K8sApplicationService type
type K8sApplicationService struct {
	mock.Mock
}

// ApplyResources provides a mock function with given fields: ctx, token, request, resourceRbacHandler
func (_m *K8sApplicationService) ApplyResources(ctx context.Context, token string, request *application.ApplyResourcesRequest, resourceRbacHandler func(string, string, k8s.ResourceRequestBean, string) bool) ([]*application.ApplyResourcesResponse, error) {
	ret := _m.Called(ctx, token, request, resourceRbacHandler)

	var r0 []*application.ApplyResourcesResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, *application.ApplyResourcesRequest, func(string, string, k8s.ResourceRequestBean, string) bool) []*application.ApplyResourcesResponse); ok {
		r0 = rf(ctx, token, request, resourceRbacHandler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*application.ApplyResourcesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *application.ApplyResourcesRequest, func(string, string, k8s.ResourceRequestBean, string) bool) error); ok {
		r1 = rf(ctx, token, request, resourceRbacHandler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePodEphemeralContainers provides a mock function with given fields: req
func (_m *K8sApplicationService) CreatePodEphemeralContainers(req *cluster.EphemeralContainerRequest) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cluster.EphemeralContainerRequest) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateResource provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) CreateResource(ctx context.Context, request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeDevtronAppId provides a mock function with given fields: applicationId
func (_m *K8sApplicationService) DecodeDevtronAppId(applicationId string) (*k8s.DevtronAppIdentifier, error) {
	ret := _m.Called(applicationId)

	var r0 *k8s.DevtronAppIdentifier
	if rf, ok := ret.Get(0).(func(string) *k8s.DevtronAppIdentifier); ok {
		r0 = rf(applicationId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.DevtronAppIdentifier)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(applicationId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResource provides a mock function with given fields: ctx, request, userId
func (_m *K8sApplicationService) DeleteResource(ctx context.Context, request *k8s.ResourceRequestBean, userId int32) (*application.ManifestResponse, error) {
	ret := _m.Called(ctx, request, userId)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean, int32) *application.ManifestResponse); ok {
		r0 = rf(ctx, request, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean, int32) error); ok {
		r1 = rf(ctx, request, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchConnectionStatusForCluster provides a mock function with given fields: k8sClientSet, clusterId
func (_m *K8sApplicationService) FetchConnectionStatusForCluster(k8sClientSet *kubernetes.Clientset, clusterId int) error {
	ret := _m.Called(k8sClientSet, clusterId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*kubernetes.Clientset, int) error); ok {
		r0 = rf(k8sClientSet, clusterId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterServiceAndIngress provides a mock function with given fields: ctx, resourceTreeInf, validRequests, appDetail, appId
func (_m *K8sApplicationService) FilterServiceAndIngress(ctx context.Context, resourceTreeInf map[string]interface{}, validRequests []k8s.ResourceRequestBean, appDetail bean.AppDetailContainer, appId string) []k8s.ResourceRequestBean {
	ret := _m.Called(ctx, resourceTreeInf, validRequests, appDetail, appId)

	var r0 []k8s.ResourceRequestBean
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, []k8s.ResourceRequestBean, bean.AppDetailContainer, string) []k8s.ResourceRequestBean); ok {
		r0 = rf(ctx, resourceTreeInf, validRequests, appDetail, appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]k8s.ResourceRequestBean)
		}
	}

	return r0
}

// GetAllApiResources provides a mock function with given fields: ctx, clusterId, isSuperAdmin, userId
func (_m *K8sApplicationService) GetAllApiResources(ctx context.Context, clusterId int, isSuperAdmin bool, userId int32) (*application.GetAllApiResourcesResponse, error) {
	ret := _m.Called(ctx, clusterId, isSuperAdmin, userId)

	var r0 *application.GetAllApiResourcesResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, bool, int32) *application.GetAllApiResourcesResponse); ok {
		r0 = rf(ctx, clusterId, isSuperAdmin, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.GetAllApiResourcesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, bool, int32) error); ok {
		r1 = rf(ctx, clusterId, isSuperAdmin, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetK8sServerVersion provides a mock function with given fields: clusterId
func (_m *K8sApplicationService) GetK8sServerVersion(clusterId int) (*version.Info, error) {
	ret := _m.Called(clusterId)

	var r0 *version.Info
	if rf, ok := ret.Get(0).(func(int) *version.Info); ok {
		r0 = rf(clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*version.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManifestsByBatch provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) GetManifestsByBatch(ctx context.Context, request []k8s.ResourceRequestBean) ([]k8s.BatchResourceResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 []k8s.BatchResourceResponse
	if rf, ok := ret.Get(0).(func(context.Context, []k8s.ResourceRequestBean) []k8s.BatchResourceResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]k8s.BatchResourceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodContainersList provides a mock function with given fields: clusterId, namespace, podName
func (_m *K8sApplicationService) GetPodContainersList(clusterId int, namespace string, podName string) (*k8s.PodContainerList, error) {
	ret := _m.Called(clusterId, namespace, podName)

	var r0 *k8s.PodContainerList
	if rf, ok := ret.Get(0).(func(int, string, string) *k8s.PodContainerList); ok {
		r0 = rf(clusterId, namespace, podName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.PodContainerList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string) error); ok {
		r1 = rf(clusterId, namespace, podName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPodLogs provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) GetPodLogs(ctx context.Context, request *k8s.ResourceRequestBean) (io.ReadCloser, error) {
	ret := _m.Called(ctx, request)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean) io.ReadCloser); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResource provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) GetResource(ctx context.Context, request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceInfo provides a mock function with given fields: ctx
func (_m *K8sApplicationService) GetResourceInfo(ctx context.Context) (*k8s.ResourceInfo, error) {
	ret := _m.Called(ctx)

	var r0 *k8s.ResourceInfo
	if rf, ok := ret.Get(0).(func(context.Context) *k8s.ResourceInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.ResourceInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceList provides a mock function with given fields: ctx, token, request, validateResourceAccess
func (_m *K8sApplicationService) GetResourceList(ctx context.Context, token string, request *k8s.ResourceRequestBean, validateResourceAccess func(string, string, k8s.ResourceRequestBean, string) bool) (*util.ClusterResourceListMap, error) {
	ret := _m.Called(ctx, token, request, validateResourceAccess)

	var r0 *util.ClusterResourceListMap
	if rf, ok := ret.Get(0).(func(context.Context, string, *k8s.ResourceRequestBean, func(string, string, k8s.ResourceRequestBean, string) bool) *util.ClusterResourceListMap); ok {
		r0 = rf(ctx, token, request, validateResourceAccess)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ClusterResourceListMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *k8s.ResourceRequestBean, func(string, string, k8s.ResourceRequestBean, string) bool) error); ok {
		r1 = rf(ctx, token, request, validateResourceAccess)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestConfigByClusterId provides a mock function with given fields: ctx, clusterId
func (_m *K8sApplicationService) GetRestConfigByClusterId(ctx context.Context, clusterId int) (*rest.Config, error) {
	ret := _m.Called(ctx, clusterId)

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func(context.Context, int) *rest.Config); ok {
		r0 = rf(ctx, clusterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, clusterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUrlsByBatch provides a mock function with given fields: ctx, resp
func (_m *K8sApplicationService) GetUrlsByBatch(ctx context.Context, resp []k8s.BatchResourceResponse) []interface{} {
	ret := _m.Called(ctx, resp)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, []k8s.BatchResourceResponse) []interface{}); ok {
		r0 = rf(ctx, resp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// ListEvents provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) ListEvents(ctx context.Context, request *k8s.ResourceRequestBean) (*application.EventsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *application.EventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean) *application.EventsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.EventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RotatePods provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) RotatePods(ctx context.Context, request *k8s.RotatePodRequest) (*k8s.RotatePodResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *k8s.RotatePodResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.RotatePodRequest) *k8s.RotatePodResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.RotatePodResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.RotatePodRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminatePodEphemeralContainer provides a mock function with given fields: req
func (_m *K8sApplicationService) TerminatePodEphemeralContainer(req cluster.EphemeralContainerRequest) (bool, error) {
	ret := _m.Called(req)

	var r0 bool
	if rf, ok := ret.Get(0).(func(cluster.EphemeralContainerRequest) bool); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cluster.EphemeralContainerRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateResource provides a mock function with given fields: ctx, request
func (_m *K8sApplicationService) UpdateResource(ctx context.Context, request *k8s.ResourceRequestBean) (*application.ManifestResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *application.ManifestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean) *application.ManifestResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*application.ManifestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateClusterResourceBean provides a mock function with given fields: ctx, clusterId, manifest, gvk, rbacCallback
func (_m *K8sApplicationService) ValidateClusterResourceBean(ctx context.Context, clusterId int, manifest unstructured.Unstructured, gvk schema.GroupVersionKind, rbacCallback func(string, application.ResourceIdentifier) bool) bool {
	ret := _m.Called(ctx, clusterId, manifest, gvk, rbacCallback)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int, unstructured.Unstructured, schema.GroupVersionKind, func(string, application.ResourceIdentifier) bool) bool); ok {
		r0 = rf(ctx, clusterId, manifest, gvk, rbacCallback)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ValidateClusterResourceRequest provides a mock function with given fields: ctx, clusterResourceRequest, rbacCallback
func (_m *K8sApplicationService) ValidateClusterResourceRequest(ctx context.Context, clusterResourceRequest *k8s.ResourceRequestBean, rbacCallback func(string, application.ResourceIdentifier) bool) (bool, error) {
	ret := _m.Called(ctx, clusterResourceRequest, rbacCallback)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *k8s.ResourceRequestBean, func(string, application.ResourceIdentifier) bool) bool); ok {
		r0 = rf(ctx, clusterResourceRequest, rbacCallback)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *k8s.ResourceRequestBean, func(string, application.ResourceIdentifier) bool) error); ok {
		r1 = rf(ctx, clusterResourceRequest, rbacCallback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatePodLogsRequestQuery provides a mock function with given fields: r
func (_m *K8sApplicationService) ValidatePodLogsRequestQuery(r *http.Request) (*k8s.ResourceRequestBean, error) {
	ret := _m.Called(r)

	var r0 *k8s.ResourceRequestBean
	if rf, ok := ret.Get(0).(func(*http.Request) *k8s.ResourceRequestBean); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*k8s.ResourceRequestBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResourceRequest provides a mock function with given fields: ctx, appIdentifier, request
func (_m *K8sApplicationService) ValidateResourceRequest(ctx context.Context, appIdentifier *client.AppIdentifier, request *application.K8sRequestBean) (bool, error) {
	ret := _m.Called(ctx, appIdentifier, request)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *client.AppIdentifier, *application.K8sRequestBean) bool); ok {
		r0 = rf(ctx, appIdentifier, request)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *client.AppIdentifier, *application.K8sRequestBean) error); ok {
		r1 = rf(ctx, appIdentifier, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateTerminalRequestQuery provides a mock function with given fields: r
func (_m *K8sApplicationService) ValidateTerminalRequestQuery(r *http.Request) (*terminal.TerminalSessionRequest, *k8s.ResourceRequestBean, error) {
	ret := _m.Called(r)

	var r0 *terminal.TerminalSessionRequest
	if rf, ok := ret.Get(0).(func(*http.Request) *terminal.TerminalSessionRequest); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*terminal.TerminalSessionRequest)
		}
	}

	var r1 *k8s.ResourceRequestBean
	if rf, ok := ret.Get(1).(func(*http.Request) *k8s.ResourceRequestBean); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*k8s.ResourceRequestBean)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*http.Request) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewK8sApplicationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewK8sApplicationService creates a new instance of K8sApplicationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewK8sApplicationService(t mockConstructorTestingTNewK8sApplicationService) *K8sApplicationService {
	mock := &K8sApplicationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
