// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/groot-windows/driver"
	"github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio/backuptar"
)

type TarStreamer struct {
	SetReaderStub        func(io.Reader)
	setReaderMutex       sync.RWMutex
	setReaderArgsForCall []struct {
		arg1 io.Reader
	}
	NextStub        func() (*tar.Header, error)
	nextMutex       sync.RWMutex
	nextArgsForCall []struct{}
	nextReturns     struct {
		result1 *tar.Header
		result2 error
	}
	nextReturnsOnCall map[int]struct {
		result1 *tar.Header
		result2 error
	}
	FileInfoFromHeaderStub        func(*tar.Header) (string, int64, *winio.FileBasicInfo, error)
	fileInfoFromHeaderMutex       sync.RWMutex
	fileInfoFromHeaderArgsForCall []struct {
		arg1 *tar.Header
	}
	fileInfoFromHeaderReturns struct {
		result1 string
		result2 int64
		result3 *winio.FileBasicInfo
		result4 error
	}
	fileInfoFromHeaderReturnsOnCall map[int]struct {
		result1 string
		result2 int64
		result3 *winio.FileBasicInfo
		result4 error
	}
	WriteBackupStreamFromTarFileStub        func(io.Writer, *tar.Header) (*tar.Header, error)
	writeBackupStreamFromTarFileMutex       sync.RWMutex
	writeBackupStreamFromTarFileArgsForCall []struct {
		arg1 io.Writer
		arg2 *tar.Header
	}
	writeBackupStreamFromTarFileReturns struct {
		result1 *tar.Header
		result2 error
	}
	writeBackupStreamFromTarFileReturnsOnCall map[int]struct {
		result1 *tar.Header
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TarStreamer) SetReader(arg1 io.Reader) {
	fake.setReaderMutex.Lock()
	fake.setReaderArgsForCall = append(fake.setReaderArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("SetReader", []interface{}{arg1})
	fake.setReaderMutex.Unlock()
	if fake.SetReaderStub != nil {
		fake.SetReaderStub(arg1)
	}
}

func (fake *TarStreamer) SetReaderCallCount() int {
	fake.setReaderMutex.RLock()
	defer fake.setReaderMutex.RUnlock()
	return len(fake.setReaderArgsForCall)
}

func (fake *TarStreamer) SetReaderArgsForCall(i int) io.Reader {
	fake.setReaderMutex.RLock()
	defer fake.setReaderMutex.RUnlock()
	return fake.setReaderArgsForCall[i].arg1
}

func (fake *TarStreamer) Next() (*tar.Header, error) {
	fake.nextMutex.Lock()
	ret, specificReturn := fake.nextReturnsOnCall[len(fake.nextArgsForCall)]
	fake.nextArgsForCall = append(fake.nextArgsForCall, struct{}{})
	fake.recordInvocation("Next", []interface{}{})
	fake.nextMutex.Unlock()
	if fake.NextStub != nil {
		return fake.NextStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.nextReturns.result1, fake.nextReturns.result2
}

func (fake *TarStreamer) NextCallCount() int {
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	return len(fake.nextArgsForCall)
}

func (fake *TarStreamer) NextReturns(result1 *tar.Header, result2 error) {
	fake.NextStub = nil
	fake.nextReturns = struct {
		result1 *tar.Header
		result2 error
	}{result1, result2}
}

func (fake *TarStreamer) NextReturnsOnCall(i int, result1 *tar.Header, result2 error) {
	fake.NextStub = nil
	if fake.nextReturnsOnCall == nil {
		fake.nextReturnsOnCall = make(map[int]struct {
			result1 *tar.Header
			result2 error
		})
	}
	fake.nextReturnsOnCall[i] = struct {
		result1 *tar.Header
		result2 error
	}{result1, result2}
}

func (fake *TarStreamer) FileInfoFromHeader(arg1 *tar.Header) (string, int64, *winio.FileBasicInfo, error) {
	fake.fileInfoFromHeaderMutex.Lock()
	ret, specificReturn := fake.fileInfoFromHeaderReturnsOnCall[len(fake.fileInfoFromHeaderArgsForCall)]
	fake.fileInfoFromHeaderArgsForCall = append(fake.fileInfoFromHeaderArgsForCall, struct {
		arg1 *tar.Header
	}{arg1})
	fake.recordInvocation("FileInfoFromHeader", []interface{}{arg1})
	fake.fileInfoFromHeaderMutex.Unlock()
	if fake.FileInfoFromHeaderStub != nil {
		return fake.FileInfoFromHeaderStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.fileInfoFromHeaderReturns.result1, fake.fileInfoFromHeaderReturns.result2, fake.fileInfoFromHeaderReturns.result3, fake.fileInfoFromHeaderReturns.result4
}

func (fake *TarStreamer) FileInfoFromHeaderCallCount() int {
	fake.fileInfoFromHeaderMutex.RLock()
	defer fake.fileInfoFromHeaderMutex.RUnlock()
	return len(fake.fileInfoFromHeaderArgsForCall)
}

func (fake *TarStreamer) FileInfoFromHeaderArgsForCall(i int) *tar.Header {
	fake.fileInfoFromHeaderMutex.RLock()
	defer fake.fileInfoFromHeaderMutex.RUnlock()
	return fake.fileInfoFromHeaderArgsForCall[i].arg1
}

func (fake *TarStreamer) FileInfoFromHeaderReturns(result1 string, result2 int64, result3 *winio.FileBasicInfo, result4 error) {
	fake.FileInfoFromHeaderStub = nil
	fake.fileInfoFromHeaderReturns = struct {
		result1 string
		result2 int64
		result3 *winio.FileBasicInfo
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *TarStreamer) FileInfoFromHeaderReturnsOnCall(i int, result1 string, result2 int64, result3 *winio.FileBasicInfo, result4 error) {
	fake.FileInfoFromHeaderStub = nil
	if fake.fileInfoFromHeaderReturnsOnCall == nil {
		fake.fileInfoFromHeaderReturnsOnCall = make(map[int]struct {
			result1 string
			result2 int64
			result3 *winio.FileBasicInfo
			result4 error
		})
	}
	fake.fileInfoFromHeaderReturnsOnCall[i] = struct {
		result1 string
		result2 int64
		result3 *winio.FileBasicInfo
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *TarStreamer) WriteBackupStreamFromTarFile(arg1 io.Writer, arg2 *tar.Header) (*tar.Header, error) {
	fake.writeBackupStreamFromTarFileMutex.Lock()
	ret, specificReturn := fake.writeBackupStreamFromTarFileReturnsOnCall[len(fake.writeBackupStreamFromTarFileArgsForCall)]
	fake.writeBackupStreamFromTarFileArgsForCall = append(fake.writeBackupStreamFromTarFileArgsForCall, struct {
		arg1 io.Writer
		arg2 *tar.Header
	}{arg1, arg2})
	fake.recordInvocation("WriteBackupStreamFromTarFile", []interface{}{arg1, arg2})
	fake.writeBackupStreamFromTarFileMutex.Unlock()
	if fake.WriteBackupStreamFromTarFileStub != nil {
		return fake.WriteBackupStreamFromTarFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeBackupStreamFromTarFileReturns.result1, fake.writeBackupStreamFromTarFileReturns.result2
}

func (fake *TarStreamer) WriteBackupStreamFromTarFileCallCount() int {
	fake.writeBackupStreamFromTarFileMutex.RLock()
	defer fake.writeBackupStreamFromTarFileMutex.RUnlock()
	return len(fake.writeBackupStreamFromTarFileArgsForCall)
}

func (fake *TarStreamer) WriteBackupStreamFromTarFileArgsForCall(i int) (io.Writer, *tar.Header) {
	fake.writeBackupStreamFromTarFileMutex.RLock()
	defer fake.writeBackupStreamFromTarFileMutex.RUnlock()
	return fake.writeBackupStreamFromTarFileArgsForCall[i].arg1, fake.writeBackupStreamFromTarFileArgsForCall[i].arg2
}

func (fake *TarStreamer) WriteBackupStreamFromTarFileReturns(result1 *tar.Header, result2 error) {
	fake.WriteBackupStreamFromTarFileStub = nil
	fake.writeBackupStreamFromTarFileReturns = struct {
		result1 *tar.Header
		result2 error
	}{result1, result2}
}

func (fake *TarStreamer) WriteBackupStreamFromTarFileReturnsOnCall(i int, result1 *tar.Header, result2 error) {
	fake.WriteBackupStreamFromTarFileStub = nil
	if fake.writeBackupStreamFromTarFileReturnsOnCall == nil {
		fake.writeBackupStreamFromTarFileReturnsOnCall = make(map[int]struct {
			result1 *tar.Header
			result2 error
		})
	}
	fake.writeBackupStreamFromTarFileReturnsOnCall[i] = struct {
		result1 *tar.Header
		result2 error
	}{result1, result2}
}

func (fake *TarStreamer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setReaderMutex.RLock()
	defer fake.setReaderMutex.RUnlock()
	fake.nextMutex.RLock()
	defer fake.nextMutex.RUnlock()
	fake.fileInfoFromHeaderMutex.RLock()
	defer fake.fileInfoFromHeaderMutex.RUnlock()
	fake.writeBackupStreamFromTarFileMutex.RLock()
	defer fake.writeBackupStreamFromTarFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TarStreamer) recordInvocation(key string, args []interface{}) {
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

var _ driver.TarStreamer = new(TarStreamer)
