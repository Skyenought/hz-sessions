// The MIT License (MIT)
//
// Copyright (c) 2016 Bo-Yi Wu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// This file may have been modified by CloudWeGo authors. All CloudWeGo
// Modifications are Copyright 2022 CloudWeGo Authors.

package redis

import (
	"testing"

	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/tester"
)

const redisTestServer = "localhost:6379"

var newRedisStore = func(_ *testing.T) sessions.Store {
	store, err := NewStore(10, "tcp", redisTestServer, "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	return store
}

func TestRedis_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newRedisStore)
}

func TestRedis_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newRedisStore)
}

func TestRedis_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newRedisStore)
}

func TestRedis_SessionClear(t *testing.T) {
	tester.Clear(t, newRedisStore)
}

func TestRedis_SessionOptions(t *testing.T) {
	tester.Options(t, newRedisStore)
}

func TestRedis_SessionMany(t *testing.T) {
	tester.Many(t, newRedisStore)
}

func TestGetRedisStore(t *testing.T) {
	t.Run("unmatched type", func(t *testing.T) {
		type store struct{ Store }
		err, rediStore := GetRedisStore(store{})
		if err == nil || rediStore != nil {
			t.Fail()
		}
	})
}