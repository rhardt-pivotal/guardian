// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/containers"
	prototypes "github.com/gogo/protobuf/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeContainer struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	InfoStub        func(context.Context) (containers.Container, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 context.Context
	}
	infoReturns struct {
		result1 containers.Container
		result2 error
	}
	infoReturnsOnCall map[int]struct {
		result1 containers.Container
		result2 error
	}
	DeleteStub        func(context.Context, ...containerd.DeleteOpts) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 []containerd.DeleteOpts
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	NewTaskStub        func(context.Context, cio.Creator, ...containerd.NewTaskOpts) (containerd.Task, error)
	newTaskMutex       sync.RWMutex
	newTaskArgsForCall []struct {
		arg1 context.Context
		arg2 cio.Creator
		arg3 []containerd.NewTaskOpts
	}
	newTaskReturns struct {
		result1 containerd.Task
		result2 error
	}
	newTaskReturnsOnCall map[int]struct {
		result1 containerd.Task
		result2 error
	}
	SpecStub        func(context.Context) (*specs.Spec, error)
	specMutex       sync.RWMutex
	specArgsForCall []struct {
		arg1 context.Context
	}
	specReturns struct {
		result1 *specs.Spec
		result2 error
	}
	specReturnsOnCall map[int]struct {
		result1 *specs.Spec
		result2 error
	}
	TaskStub        func(context.Context, cio.Attach) (containerd.Task, error)
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1 context.Context
		arg2 cio.Attach
	}
	taskReturns struct {
		result1 containerd.Task
		result2 error
	}
	taskReturnsOnCall map[int]struct {
		result1 containerd.Task
		result2 error
	}
	ImageStub        func(context.Context) (containerd.Image, error)
	imageMutex       sync.RWMutex
	imageArgsForCall []struct {
		arg1 context.Context
	}
	imageReturns struct {
		result1 containerd.Image
		result2 error
	}
	imageReturnsOnCall map[int]struct {
		result1 containerd.Image
		result2 error
	}
	LabelsStub        func(context.Context) (map[string]string, error)
	labelsMutex       sync.RWMutex
	labelsArgsForCall []struct {
		arg1 context.Context
	}
	labelsReturns struct {
		result1 map[string]string
		result2 error
	}
	labelsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	SetLabelsStub        func(context.Context, map[string]string) (map[string]string, error)
	setLabelsMutex       sync.RWMutex
	setLabelsArgsForCall []struct {
		arg1 context.Context
		arg2 map[string]string
	}
	setLabelsReturns struct {
		result1 map[string]string
		result2 error
	}
	setLabelsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ExtensionsStub        func(context.Context) (map[string]prototypes.Any, error)
	extensionsMutex       sync.RWMutex
	extensionsArgsForCall []struct {
		arg1 context.Context
	}
	extensionsReturns struct {
		result1 map[string]prototypes.Any
		result2 error
	}
	extensionsReturnsOnCall map[int]struct {
		result1 map[string]prototypes.Any
		result2 error
	}
	UpdateStub        func(context.Context, ...containerd.UpdateContainerOpts) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 []containerd.UpdateContainerOpts
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainer) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeContainer) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeContainer) IDReturns(result1 string) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContainer) IDReturnsOnCall(i int, result1 string) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeContainer) Info(arg1 context.Context) (containers.Container, error) {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Info", []interface{}{arg1})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.infoReturns.result1, fake.infoReturns.result2
}

func (fake *FakeContainer) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeContainer) InfoArgsForCall(i int) context.Context {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].arg1
}

func (fake *FakeContainer) InfoReturns(result1 containers.Container, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 containers.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) InfoReturnsOnCall(i int, result1 containers.Container, result2 error) {
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 containers.Container
			result2 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 containers.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Delete(arg1 context.Context, arg2 ...containerd.DeleteOpts) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 []containerd.DeleteOpts
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeContainer) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeContainer) DeleteArgsForCall(i int) (context.Context, []containerd.DeleteOpts) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1, fake.deleteArgsForCall[i].arg2
}

func (fake *FakeContainer) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) DeleteReturnsOnCall(i int, result1 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) NewTask(arg1 context.Context, arg2 cio.Creator, arg3 ...containerd.NewTaskOpts) (containerd.Task, error) {
	fake.newTaskMutex.Lock()
	ret, specificReturn := fake.newTaskReturnsOnCall[len(fake.newTaskArgsForCall)]
	fake.newTaskArgsForCall = append(fake.newTaskArgsForCall, struct {
		arg1 context.Context
		arg2 cio.Creator
		arg3 []containerd.NewTaskOpts
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewTask", []interface{}{arg1, arg2, arg3})
	fake.newTaskMutex.Unlock()
	if fake.NewTaskStub != nil {
		return fake.NewTaskStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.newTaskReturns.result1, fake.newTaskReturns.result2
}

func (fake *FakeContainer) NewTaskCallCount() int {
	fake.newTaskMutex.RLock()
	defer fake.newTaskMutex.RUnlock()
	return len(fake.newTaskArgsForCall)
}

func (fake *FakeContainer) NewTaskArgsForCall(i int) (context.Context, cio.Creator, []containerd.NewTaskOpts) {
	fake.newTaskMutex.RLock()
	defer fake.newTaskMutex.RUnlock()
	return fake.newTaskArgsForCall[i].arg1, fake.newTaskArgsForCall[i].arg2, fake.newTaskArgsForCall[i].arg3
}

func (fake *FakeContainer) NewTaskReturns(result1 containerd.Task, result2 error) {
	fake.NewTaskStub = nil
	fake.newTaskReturns = struct {
		result1 containerd.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) NewTaskReturnsOnCall(i int, result1 containerd.Task, result2 error) {
	fake.NewTaskStub = nil
	if fake.newTaskReturnsOnCall == nil {
		fake.newTaskReturnsOnCall = make(map[int]struct {
			result1 containerd.Task
			result2 error
		})
	}
	fake.newTaskReturnsOnCall[i] = struct {
		result1 containerd.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Spec(arg1 context.Context) (*specs.Spec, error) {
	fake.specMutex.Lock()
	ret, specificReturn := fake.specReturnsOnCall[len(fake.specArgsForCall)]
	fake.specArgsForCall = append(fake.specArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Spec", []interface{}{arg1})
	fake.specMutex.Unlock()
	if fake.SpecStub != nil {
		return fake.SpecStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.specReturns.result1, fake.specReturns.result2
}

func (fake *FakeContainer) SpecCallCount() int {
	fake.specMutex.RLock()
	defer fake.specMutex.RUnlock()
	return len(fake.specArgsForCall)
}

func (fake *FakeContainer) SpecArgsForCall(i int) context.Context {
	fake.specMutex.RLock()
	defer fake.specMutex.RUnlock()
	return fake.specArgsForCall[i].arg1
}

func (fake *FakeContainer) SpecReturns(result1 *specs.Spec, result2 error) {
	fake.SpecStub = nil
	fake.specReturns = struct {
		result1 *specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) SpecReturnsOnCall(i int, result1 *specs.Spec, result2 error) {
	fake.SpecStub = nil
	if fake.specReturnsOnCall == nil {
		fake.specReturnsOnCall = make(map[int]struct {
			result1 *specs.Spec
			result2 error
		})
	}
	fake.specReturnsOnCall[i] = struct {
		result1 *specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Task(arg1 context.Context, arg2 cio.Attach) (containerd.Task, error) {
	fake.taskMutex.Lock()
	ret, specificReturn := fake.taskReturnsOnCall[len(fake.taskArgsForCall)]
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1 context.Context
		arg2 cio.Attach
	}{arg1, arg2})
	fake.recordInvocation("Task", []interface{}{arg1, arg2})
	fake.taskMutex.Unlock()
	if fake.TaskStub != nil {
		return fake.TaskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.taskReturns.result1, fake.taskReturns.result2
}

func (fake *FakeContainer) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeContainer) TaskArgsForCall(i int) (context.Context, cio.Attach) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.taskArgsForCall[i].arg1, fake.taskArgsForCall[i].arg2
}

func (fake *FakeContainer) TaskReturns(result1 containerd.Task, result2 error) {
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 containerd.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) TaskReturnsOnCall(i int, result1 containerd.Task, result2 error) {
	fake.TaskStub = nil
	if fake.taskReturnsOnCall == nil {
		fake.taskReturnsOnCall = make(map[int]struct {
			result1 containerd.Task
			result2 error
		})
	}
	fake.taskReturnsOnCall[i] = struct {
		result1 containerd.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Image(arg1 context.Context) (containerd.Image, error) {
	fake.imageMutex.Lock()
	ret, specificReturn := fake.imageReturnsOnCall[len(fake.imageArgsForCall)]
	fake.imageArgsForCall = append(fake.imageArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Image", []interface{}{arg1})
	fake.imageMutex.Unlock()
	if fake.ImageStub != nil {
		return fake.ImageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.imageReturns.result1, fake.imageReturns.result2
}

func (fake *FakeContainer) ImageCallCount() int {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return len(fake.imageArgsForCall)
}

func (fake *FakeContainer) ImageArgsForCall(i int) context.Context {
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	return fake.imageArgsForCall[i].arg1
}

func (fake *FakeContainer) ImageReturns(result1 containerd.Image, result2 error) {
	fake.ImageStub = nil
	fake.imageReturns = struct {
		result1 containerd.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) ImageReturnsOnCall(i int, result1 containerd.Image, result2 error) {
	fake.ImageStub = nil
	if fake.imageReturnsOnCall == nil {
		fake.imageReturnsOnCall = make(map[int]struct {
			result1 containerd.Image
			result2 error
		})
	}
	fake.imageReturnsOnCall[i] = struct {
		result1 containerd.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Labels(arg1 context.Context) (map[string]string, error) {
	fake.labelsMutex.Lock()
	ret, specificReturn := fake.labelsReturnsOnCall[len(fake.labelsArgsForCall)]
	fake.labelsArgsForCall = append(fake.labelsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Labels", []interface{}{arg1})
	fake.labelsMutex.Unlock()
	if fake.LabelsStub != nil {
		return fake.LabelsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.labelsReturns.result1, fake.labelsReturns.result2
}

func (fake *FakeContainer) LabelsCallCount() int {
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	return len(fake.labelsArgsForCall)
}

func (fake *FakeContainer) LabelsArgsForCall(i int) context.Context {
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	return fake.labelsArgsForCall[i].arg1
}

func (fake *FakeContainer) LabelsReturns(result1 map[string]string, result2 error) {
	fake.LabelsStub = nil
	fake.labelsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LabelsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.LabelsStub = nil
	if fake.labelsReturnsOnCall == nil {
		fake.labelsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.labelsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) SetLabels(arg1 context.Context, arg2 map[string]string) (map[string]string, error) {
	fake.setLabelsMutex.Lock()
	ret, specificReturn := fake.setLabelsReturnsOnCall[len(fake.setLabelsArgsForCall)]
	fake.setLabelsArgsForCall = append(fake.setLabelsArgsForCall, struct {
		arg1 context.Context
		arg2 map[string]string
	}{arg1, arg2})
	fake.recordInvocation("SetLabels", []interface{}{arg1, arg2})
	fake.setLabelsMutex.Unlock()
	if fake.SetLabelsStub != nil {
		return fake.SetLabelsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setLabelsReturns.result1, fake.setLabelsReturns.result2
}

func (fake *FakeContainer) SetLabelsCallCount() int {
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	return len(fake.setLabelsArgsForCall)
}

func (fake *FakeContainer) SetLabelsArgsForCall(i int) (context.Context, map[string]string) {
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	return fake.setLabelsArgsForCall[i].arg1, fake.setLabelsArgsForCall[i].arg2
}

func (fake *FakeContainer) SetLabelsReturns(result1 map[string]string, result2 error) {
	fake.SetLabelsStub = nil
	fake.setLabelsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) SetLabelsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.SetLabelsStub = nil
	if fake.setLabelsReturnsOnCall == nil {
		fake.setLabelsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.setLabelsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Extensions(arg1 context.Context) (map[string]prototypes.Any, error) {
	fake.extensionsMutex.Lock()
	ret, specificReturn := fake.extensionsReturnsOnCall[len(fake.extensionsArgsForCall)]
	fake.extensionsArgsForCall = append(fake.extensionsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Extensions", []interface{}{arg1})
	fake.extensionsMutex.Unlock()
	if fake.ExtensionsStub != nil {
		return fake.ExtensionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.extensionsReturns.result1, fake.extensionsReturns.result2
}

func (fake *FakeContainer) ExtensionsCallCount() int {
	fake.extensionsMutex.RLock()
	defer fake.extensionsMutex.RUnlock()
	return len(fake.extensionsArgsForCall)
}

func (fake *FakeContainer) ExtensionsArgsForCall(i int) context.Context {
	fake.extensionsMutex.RLock()
	defer fake.extensionsMutex.RUnlock()
	return fake.extensionsArgsForCall[i].arg1
}

func (fake *FakeContainer) ExtensionsReturns(result1 map[string]prototypes.Any, result2 error) {
	fake.ExtensionsStub = nil
	fake.extensionsReturns = struct {
		result1 map[string]prototypes.Any
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) ExtensionsReturnsOnCall(i int, result1 map[string]prototypes.Any, result2 error) {
	fake.ExtensionsStub = nil
	if fake.extensionsReturnsOnCall == nil {
		fake.extensionsReturnsOnCall = make(map[int]struct {
			result1 map[string]prototypes.Any
			result2 error
		})
	}
	fake.extensionsReturnsOnCall[i] = struct {
		result1 map[string]prototypes.Any
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Update(arg1 context.Context, arg2 ...containerd.UpdateContainerOpts) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 []containerd.UpdateContainerOpts
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateReturns.result1
}

func (fake *FakeContainer) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeContainer) UpdateArgsForCall(i int) (context.Context, []containerd.UpdateContainerOpts) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeContainer) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) UpdateReturnsOnCall(i int, result1 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.newTaskMutex.RLock()
	defer fake.newTaskMutex.RUnlock()
	fake.specMutex.RLock()
	defer fake.specMutex.RUnlock()
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	fake.imageMutex.RLock()
	defer fake.imageMutex.RUnlock()
	fake.labelsMutex.RLock()
	defer fake.labelsMutex.RUnlock()
	fake.setLabelsMutex.RLock()
	defer fake.setLabelsMutex.RUnlock()
	fake.extensionsMutex.RLock()
	defer fake.extensionsMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runcontainerd.Container = new(FakeContainer)