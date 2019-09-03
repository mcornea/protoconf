package libprotoconf

import (
	"fmt"

	protoconfvalue "protoconf.com/datatypes/proto/v1/protoconfvalue"
)

type TestWatcher struct {
	valuesChan chan *protoconfvalue.ProtoconfValue
}

// NewTestWatcher creates a new file-backed protoconf watcher
func NewTestWatcher(ch chan *protoconfvalue.ProtoconfValue) (Watcher, error) {
	watcher := &TestWatcher{
		valuesChan: ch,
	}

	return watcher, nil
}

// Watch a value given its path
func (w *TestWatcher) Watch(path string, stopCh <-chan struct{}) (<-chan Result, error) {
	watchCh := make(chan Result)
	go func() {
		defer func() {
			close(watchCh)
		}()

		for {
			select {
			case protoconfValue := <-w.valuesChan:

				any, err := injectSecrets(protoconfValue)
				if err != nil {
					watchCh <- Result{nil, fmt.Errorf("error injecting secrets path=%s protoconf_value=%s err=%s", path, protoconfValue, err)}
					return
				}
				watchCh <- Result{any, nil}

				select {
				case <-stopCh:
					return
				}
			}
		}
	}()

	return watchCh, nil
}

func (w *TestWatcher) Close() {
}
