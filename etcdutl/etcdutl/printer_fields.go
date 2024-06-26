// Copyright 2021 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package etcdutl

import (
	"fmt"

	"go.etcd.io/etcd/etcdutl/v3/snapshot"
)

type fieldsPrinter struct{ printer }

func (p *fieldsPrinter) DBStatus(r snapshot.Status) {
	fmt.Println(`"Hash" :`, r.Hash)
	fmt.Println(`"Revision" :`, r.Revision)
	fmt.Println(`"Keys" :`, r.TotalKey)
	fmt.Println(`"Size" :`, r.TotalSize)
	fmt.Println(`"Version" :`, r.Version)
}

func (p *fieldsPrinter) DBHashKV(r HashKV) {
	fmt.Println(`"Hash" :`, r.Hash)
	fmt.Println(`"Hash revision" :`, r.HashRevision)
	fmt.Println(`"Compact revision" :`, r.CompactRevision)
}
