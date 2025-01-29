// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/controller-manager-library/pkg/resources (interfaces: Object)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	resources "github.com/gardener/controller-manager-library/pkg/resources"
	abstract "github.com/gardener/controller-manager-library/pkg/resources/abstract"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
)

// MockObject is a mock of Object interface.
type MockObject struct {
	ctrl     *gomock.Controller
	recorder *MockObjectMockRecorder
}

// MockObjectMockRecorder is the mock recorder for MockObject.
type MockObjectMockRecorder struct {
	mock *MockObject
}

// NewMockObject creates a new mock instance.
func NewMockObject(ctrl *gomock.Controller) *MockObject {
	mock := &MockObject{ctrl: ctrl}
	mock.recorder = &MockObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObject) EXPECT() *MockObjectMockRecorder {
	return m.recorder
}

// AddOwner mocks base method.
func (m *MockObject) AddOwner(arg0 resources.Object) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOwner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddOwner indicates an expected call of AddOwner.
func (mr *MockObjectMockRecorder) AddOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOwner", reflect.TypeOf((*MockObject)(nil).AddOwner), arg0)
}

// AnnotatedEventf mocks base method.
func (m *MockObject) AnnotatedEventf(arg0 map[string]string, arg1, arg2, arg3 string, arg4 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AnnotatedEventf", varargs...)
}

// AnnotatedEventf indicates an expected call of AnnotatedEventf.
func (mr *MockObjectMockRecorder) AnnotatedEventf(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotatedEventf", reflect.TypeOf((*MockObject)(nil).AnnotatedEventf), varargs...)
}

// ClusterKey mocks base method.
func (m *MockObject) ClusterKey() abstract.ClusterObjectKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterKey")
	ret0, _ := ret[0].(abstract.ClusterObjectKey)
	return ret0
}

// ClusterKey indicates an expected call of ClusterKey.
func (mr *MockObjectMockRecorder) ClusterKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterKey", reflect.TypeOf((*MockObject)(nil).ClusterKey))
}

// Create mocks base method.
func (m *MockObject) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockObjectMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockObject)(nil).Create))
}

// CreateOrModify mocks base method.
func (m *MockObject) CreateOrModify(arg0 resources.Modifier) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrModify", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrModify indicates an expected call of CreateOrModify.
func (mr *MockObjectMockRecorder) CreateOrModify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrModify", reflect.TypeOf((*MockObject)(nil).CreateOrModify), arg0)
}

// CreateOrUpdate mocks base method.
func (m *MockObject) CreateOrUpdate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockObjectMockRecorder) CreateOrUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockObject)(nil).CreateOrUpdate))
}

// Data mocks base method.
func (m *MockObject) Data() abstract.ObjectData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(abstract.ObjectData)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockObjectMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockObject)(nil).Data))
}

// DeepCopy mocks base method.
func (m *MockObject) DeepCopy() resources.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeepCopy")
	ret0, _ := ret[0].(resources.Object)
	return ret0
}

// DeepCopy indicates an expected call of DeepCopy.
func (mr *MockObjectMockRecorder) DeepCopy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopy", reflect.TypeOf((*MockObject)(nil).DeepCopy))
}

// Delete mocks base method.
func (m *MockObject) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockObjectMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockObject)(nil).Delete))
}

// Description mocks base method.
func (m *MockObject) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockObjectMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockObject)(nil).Description))
}

// Event mocks base method.
func (m *MockObject) Event(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Event", arg0, arg1, arg2)
}

// Event indicates an expected call of Event.
func (mr *MockObjectMockRecorder) Event(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Event", reflect.TypeOf((*MockObject)(nil).Event), arg0, arg1, arg2)
}

// Eventf mocks base method.
func (m *MockObject) Eventf(arg0, arg1, arg2 string, arg3 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Eventf", varargs...)
}

// Eventf indicates an expected call of Eventf.
func (mr *MockObjectMockRecorder) Eventf(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eventf", reflect.TypeOf((*MockObject)(nil).Eventf), varargs...)
}

// GetAnnotation mocks base method.
func (m *MockObject) GetAnnotation(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotation", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAnnotation indicates an expected call of GetAnnotation.
func (mr *MockObjectMockRecorder) GetAnnotation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotation", reflect.TypeOf((*MockObject)(nil).GetAnnotation), arg0)
}

// GetAnnotations mocks base method.
func (m *MockObject) GetAnnotations() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockObjectMockRecorder) GetAnnotations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockObject)(nil).GetAnnotations))
}

// GetCluster mocks base method.
func (m *MockObject) GetCluster() resources.Cluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster")
	ret0, _ := ret[0].(resources.Cluster)
	return ret0
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockObjectMockRecorder) GetCluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockObject)(nil).GetCluster))
}

// GetCreationTimestamp mocks base method.
func (m *MockObject) GetCreationTimestamp() v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreationTimestamp")
	ret0, _ := ret[0].(v1.Time)
	return ret0
}

// GetCreationTimestamp indicates an expected call of GetCreationTimestamp.
func (mr *MockObjectMockRecorder) GetCreationTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).GetCreationTimestamp))
}

// GetDeletionGracePeriodSeconds mocks base method.
func (m *MockObject) GetDeletionGracePeriodSeconds() *int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionGracePeriodSeconds")
	ret0, _ := ret[0].(*int64)
	return ret0
}

// GetDeletionGracePeriodSeconds indicates an expected call of GetDeletionGracePeriodSeconds.
func (mr *MockObjectMockRecorder) GetDeletionGracePeriodSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).GetDeletionGracePeriodSeconds))
}

// GetDeletionTimestamp mocks base method.
func (m *MockObject) GetDeletionTimestamp() *v1.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletionTimestamp")
	ret0, _ := ret[0].(*v1.Time)
	return ret0
}

// GetDeletionTimestamp indicates an expected call of GetDeletionTimestamp.
func (mr *MockObjectMockRecorder) GetDeletionTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).GetDeletionTimestamp))
}

// GetFinalizers mocks base method.
func (m *MockObject) GetFinalizers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetFinalizers indicates an expected call of GetFinalizers.
func (mr *MockObjectMockRecorder) GetFinalizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizers", reflect.TypeOf((*MockObject)(nil).GetFinalizers))
}

// GetFullObject mocks base method.
func (m *MockObject) GetFullObject() (resources.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullObject")
	ret0, _ := ret[0].(resources.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullObject indicates an expected call of GetFullObject.
func (mr *MockObjectMockRecorder) GetFullObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullObject", reflect.TypeOf((*MockObject)(nil).GetFullObject))
}

// GetGenerateName mocks base method.
func (m *MockObject) GetGenerateName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerateName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGenerateName indicates an expected call of GetGenerateName.
func (mr *MockObjectMockRecorder) GetGenerateName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerateName", reflect.TypeOf((*MockObject)(nil).GetGenerateName))
}

// GetGeneration mocks base method.
func (m *MockObject) GetGeneration() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGeneration")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetGeneration indicates an expected call of GetGeneration.
func (mr *MockObjectMockRecorder) GetGeneration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGeneration", reflect.TypeOf((*MockObject)(nil).GetGeneration))
}

// GetLabel mocks base method.
func (m *MockObject) GetLabel(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabel", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLabel indicates an expected call of GetLabel.
func (mr *MockObjectMockRecorder) GetLabel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabel", reflect.TypeOf((*MockObject)(nil).GetLabel), arg0)
}

// GetLabels mocks base method.
func (m *MockObject) GetLabels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockObjectMockRecorder) GetLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockObject)(nil).GetLabels))
}

// GetManagedFields mocks base method.
func (m *MockObject) GetManagedFields() []v1.ManagedFieldsEntry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedFields")
	ret0, _ := ret[0].([]v1.ManagedFieldsEntry)
	return ret0
}

// GetManagedFields indicates an expected call of GetManagedFields.
func (mr *MockObjectMockRecorder) GetManagedFields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedFields", reflect.TypeOf((*MockObject)(nil).GetManagedFields))
}

// GetName mocks base method.
func (m *MockObject) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockObjectMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockObject)(nil).GetName))
}

// GetNamespace mocks base method.
func (m *MockObject) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockObjectMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockObject)(nil).GetNamespace))
}

// GetOwnerReference mocks base method.
func (m *MockObject) GetOwnerReference() *v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReference")
	ret0, _ := ret[0].(*v1.OwnerReference)
	return ret0
}

// GetOwnerReference indicates an expected call of GetOwnerReference.
func (mr *MockObjectMockRecorder) GetOwnerReference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReference", reflect.TypeOf((*MockObject)(nil).GetOwnerReference))
}

// GetOwnerReferences mocks base method.
func (m *MockObject) GetOwnerReferences() []v1.OwnerReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerReferences")
	ret0, _ := ret[0].([]v1.OwnerReference)
	return ret0
}

// GetOwnerReferences indicates an expected call of GetOwnerReferences.
func (mr *MockObjectMockRecorder) GetOwnerReferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerReferences", reflect.TypeOf((*MockObject)(nil).GetOwnerReferences))
}

// GetOwners mocks base method.
func (m *MockObject) GetOwners(arg0 ...schema.GroupKind) abstract.ClusterObjectKeySet {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOwners", varargs...)
	ret0, _ := ret[0].(abstract.ClusterObjectKeySet)
	return ret0
}

// GetOwners indicates an expected call of GetOwners.
func (mr *MockObjectMockRecorder) GetOwners(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwners", reflect.TypeOf((*MockObject)(nil).GetOwners), arg0...)
}

// GetResource mocks base method.
func (m *MockObject) GetResource() resources.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource")
	ret0, _ := ret[0].(resources.Interface)
	return ret0
}

// GetResource indicates an expected call of GetResource.
func (mr *MockObjectMockRecorder) GetResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockObject)(nil).GetResource))
}

// GetResourceVersion mocks base method.
func (m *MockObject) GetResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResourceVersion indicates an expected call of GetResourceVersion.
func (mr *MockObjectMockRecorder) GetResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceVersion", reflect.TypeOf((*MockObject)(nil).GetResourceVersion))
}

// GetSelfLink mocks base method.
func (m *MockObject) GetSelfLink() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfLink")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSelfLink indicates an expected call of GetSelfLink.
func (mr *MockObjectMockRecorder) GetSelfLink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfLink", reflect.TypeOf((*MockObject)(nil).GetSelfLink))
}

// GetUID mocks base method.
func (m *MockObject) GetUID() types.UID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUID")
	ret0, _ := ret[0].(types.UID)
	return ret0
}

// GetUID indicates an expected call of GetUID.
func (mr *MockObjectMockRecorder) GetUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUID", reflect.TypeOf((*MockObject)(nil).GetUID))
}

// GroupKind mocks base method.
func (m *MockObject) GroupKind() schema.GroupKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupKind")
	ret0, _ := ret[0].(schema.GroupKind)
	return ret0
}

// GroupKind indicates an expected call of GroupKind.
func (mr *MockObjectMockRecorder) GroupKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupKind", reflect.TypeOf((*MockObject)(nil).GroupKind))
}

// GroupVersionKind mocks base method.
func (m *MockObject) GroupVersionKind() schema.GroupVersionKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupVersionKind")
	ret0, _ := ret[0].(schema.GroupVersionKind)
	return ret0
}

// GroupVersionKind indicates an expected call of GroupVersionKind.
func (mr *MockObjectMockRecorder) GroupVersionKind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupVersionKind", reflect.TypeOf((*MockObject)(nil).GroupVersionKind))
}

// HasFinalizer mocks base method.
func (m *MockObject) HasFinalizer(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasFinalizer", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasFinalizer indicates an expected call of HasFinalizer.
func (mr *MockObjectMockRecorder) HasFinalizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasFinalizer", reflect.TypeOf((*MockObject)(nil).HasFinalizer), arg0)
}

// IsA mocks base method.
func (m *MockObject) IsA(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsA", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsA indicates an expected call of IsA.
func (mr *MockObjectMockRecorder) IsA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsA", reflect.TypeOf((*MockObject)(nil).IsA), arg0)
}

// IsCoLocatedTo mocks base method.
func (m *MockObject) IsCoLocatedTo(arg0 resources.Object) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCoLocatedTo", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCoLocatedTo indicates an expected call of IsCoLocatedTo.
func (mr *MockObjectMockRecorder) IsCoLocatedTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCoLocatedTo", reflect.TypeOf((*MockObject)(nil).IsCoLocatedTo), arg0)
}

// IsDeleting mocks base method.
func (m *MockObject) IsDeleting() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDeleting")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDeleting indicates an expected call of IsDeleting.
func (mr *MockObjectMockRecorder) IsDeleting() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDeleting", reflect.TypeOf((*MockObject)(nil).IsDeleting))
}

// IsMinimal mocks base method.
func (m *MockObject) IsMinimal() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMinimal")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMinimal indicates an expected call of IsMinimal.
func (mr *MockObjectMockRecorder) IsMinimal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMinimal", reflect.TypeOf((*MockObject)(nil).IsMinimal))
}

// Key mocks base method.
func (m *MockObject) Key() abstract.ObjectKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(abstract.ObjectKey)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockObjectMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockObject)(nil).Key))
}

// MinimalData mocks base method.
func (m *MockObject) MinimalData() *v1.PartialObjectMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinimalData")
	ret0, _ := ret[0].(*v1.PartialObjectMetadata)
	return ret0
}

// MinimalData indicates an expected call of MinimalData.
func (mr *MockObjectMockRecorder) MinimalData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinimalData", reflect.TypeOf((*MockObject)(nil).MinimalData))
}

// Modify mocks base method.
func (m *MockObject) Modify(arg0 resources.Modifier) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Modify indicates an expected call of Modify.
func (mr *MockObjectMockRecorder) Modify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockObject)(nil).Modify), arg0)
}

// ModifyStatus mocks base method.
func (m *MockObject) ModifyStatus(arg0 resources.Modifier) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyStatus", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyStatus indicates an expected call of ModifyStatus.
func (mr *MockObjectMockRecorder) ModifyStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyStatus", reflect.TypeOf((*MockObject)(nil).ModifyStatus), arg0)
}

// ObjectName mocks base method.
func (m *MockObject) ObjectName() abstract.ObjectName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectName")
	ret0, _ := ret[0].(abstract.ObjectName)
	return ret0
}

// ObjectName indicates an expected call of ObjectName.
func (mr *MockObjectMockRecorder) ObjectName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectName", reflect.TypeOf((*MockObject)(nil).ObjectName))
}

// RemoveFinalizer mocks base method.
func (m *MockObject) RemoveFinalizer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFinalizer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFinalizer indicates an expected call of RemoveFinalizer.
func (mr *MockObjectMockRecorder) RemoveFinalizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizer", reflect.TypeOf((*MockObject)(nil).RemoveFinalizer), arg0)
}

// RemoveOwner mocks base method.
func (m *MockObject) RemoveOwner(arg0 resources.Object) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOwner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemoveOwner indicates an expected call of RemoveOwner.
func (mr *MockObjectMockRecorder) RemoveOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOwner", reflect.TypeOf((*MockObject)(nil).RemoveOwner), arg0)
}

// Resources mocks base method.
func (m *MockObject) Resources() resources.Resources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].(resources.Resources)
	return ret0
}

// Resources indicates an expected call of Resources.
func (mr *MockObjectMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockObject)(nil).Resources))
}

// SetAnnotations mocks base method.
func (m *MockObject) SetAnnotations(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotations", arg0)
}

// SetAnnotations indicates an expected call of SetAnnotations.
func (mr *MockObjectMockRecorder) SetAnnotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotations", reflect.TypeOf((*MockObject)(nil).SetAnnotations), arg0)
}

// SetCreationTimestamp mocks base method.
func (m *MockObject) SetCreationTimestamp(arg0 v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCreationTimestamp", arg0)
}

// SetCreationTimestamp indicates an expected call of SetCreationTimestamp.
func (mr *MockObjectMockRecorder) SetCreationTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCreationTimestamp", reflect.TypeOf((*MockObject)(nil).SetCreationTimestamp), arg0)
}

// SetDeletionGracePeriodSeconds mocks base method.
func (m *MockObject) SetDeletionGracePeriodSeconds(arg0 *int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionGracePeriodSeconds", arg0)
}

// SetDeletionGracePeriodSeconds indicates an expected call of SetDeletionGracePeriodSeconds.
func (mr *MockObjectMockRecorder) SetDeletionGracePeriodSeconds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionGracePeriodSeconds", reflect.TypeOf((*MockObject)(nil).SetDeletionGracePeriodSeconds), arg0)
}

// SetDeletionTimestamp mocks base method.
func (m *MockObject) SetDeletionTimestamp(arg0 *v1.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDeletionTimestamp", arg0)
}

// SetDeletionTimestamp indicates an expected call of SetDeletionTimestamp.
func (mr *MockObjectMockRecorder) SetDeletionTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeletionTimestamp", reflect.TypeOf((*MockObject)(nil).SetDeletionTimestamp), arg0)
}

// SetFinalizer mocks base method.
func (m *MockObject) SetFinalizer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFinalizer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFinalizer indicates an expected call of SetFinalizer.
func (mr *MockObjectMockRecorder) SetFinalizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizer", reflect.TypeOf((*MockObject)(nil).SetFinalizer), arg0)
}

// SetFinalizers mocks base method.
func (m *MockObject) SetFinalizers(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalizers", arg0)
}

// SetFinalizers indicates an expected call of SetFinalizers.
func (mr *MockObjectMockRecorder) SetFinalizers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalizers", reflect.TypeOf((*MockObject)(nil).SetFinalizers), arg0)
}

// SetGenerateName mocks base method.
func (m *MockObject) SetGenerateName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGenerateName", arg0)
}

// SetGenerateName indicates an expected call of SetGenerateName.
func (mr *MockObjectMockRecorder) SetGenerateName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGenerateName", reflect.TypeOf((*MockObject)(nil).SetGenerateName), arg0)
}

// SetGeneration mocks base method.
func (m *MockObject) SetGeneration(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGeneration", arg0)
}

// SetGeneration indicates an expected call of SetGeneration.
func (mr *MockObjectMockRecorder) SetGeneration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGeneration", reflect.TypeOf((*MockObject)(nil).SetGeneration), arg0)
}

// SetLabels mocks base method.
func (m *MockObject) SetLabels(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLabels", arg0)
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockObjectMockRecorder) SetLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockObject)(nil).SetLabels), arg0)
}

// SetManagedFields mocks base method.
func (m *MockObject) SetManagedFields(arg0 []v1.ManagedFieldsEntry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetManagedFields", arg0)
}

// SetManagedFields indicates an expected call of SetManagedFields.
func (mr *MockObjectMockRecorder) SetManagedFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetManagedFields", reflect.TypeOf((*MockObject)(nil).SetManagedFields), arg0)
}

// SetName mocks base method.
func (m *MockObject) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockObjectMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockObject)(nil).SetName), arg0)
}

// SetNamespace mocks base method.
func (m *MockObject) SetNamespace(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespace", arg0)
}

// SetNamespace indicates an expected call of SetNamespace.
func (mr *MockObjectMockRecorder) SetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespace", reflect.TypeOf((*MockObject)(nil).SetNamespace), arg0)
}

// SetOwnerReferences mocks base method.
func (m *MockObject) SetOwnerReferences(arg0 []v1.OwnerReference) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOwnerReferences", arg0)
}

// SetOwnerReferences indicates an expected call of SetOwnerReferences.
func (mr *MockObjectMockRecorder) SetOwnerReferences(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerReferences", reflect.TypeOf((*MockObject)(nil).SetOwnerReferences), arg0)
}

// SetResourceVersion mocks base method.
func (m *MockObject) SetResourceVersion(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResourceVersion", arg0)
}

// SetResourceVersion indicates an expected call of SetResourceVersion.
func (mr *MockObjectMockRecorder) SetResourceVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResourceVersion", reflect.TypeOf((*MockObject)(nil).SetResourceVersion), arg0)
}

// SetSelfLink mocks base method.
func (m *MockObject) SetSelfLink(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSelfLink", arg0)
}

// SetSelfLink indicates an expected call of SetSelfLink.
func (mr *MockObjectMockRecorder) SetSelfLink(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSelfLink", reflect.TypeOf((*MockObject)(nil).SetSelfLink), arg0)
}

// SetUID mocks base method.
func (m *MockObject) SetUID(arg0 types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUID", arg0)
}

// SetUID indicates an expected call of SetUID.
func (mr *MockObjectMockRecorder) SetUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUID", reflect.TypeOf((*MockObject)(nil).SetUID), arg0)
}

// StatusField mocks base method.
func (m *MockObject) StatusField() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusField")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// StatusField indicates an expected call of StatusField.
func (mr *MockObjectMockRecorder) StatusField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusField", reflect.TypeOf((*MockObject)(nil).StatusField))
}

// Update mocks base method.
func (m *MockObject) Update() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update")
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockObjectMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockObject)(nil).Update))
}

// UpdateFromCache mocks base method.
func (m *MockObject) UpdateFromCache() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFromCache")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFromCache indicates an expected call of UpdateFromCache.
func (mr *MockObjectMockRecorder) UpdateFromCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFromCache", reflect.TypeOf((*MockObject)(nil).UpdateFromCache))
}

// UpdateStatus mocks base method.
func (m *MockObject) UpdateStatus() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus")
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockObjectMockRecorder) UpdateStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockObject)(nil).UpdateStatus))
}
