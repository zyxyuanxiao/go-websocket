package servers

import (
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func newTestManager() {

}

func TestAddClient(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)

	manager.AddClient(clientSocket)

	Convey("测试添加客户端", t, func() {
		Convey("长度是否够", func() {
			So(len(manager.ClientIdMap), ShouldEqual, 1)
		})

		Convey("clientId是否存在", func() {
			_, ok := manager.ClientIdMap[clientId]
			So(ok, ShouldBeTrue)
		})
	})
}

func TestDelClient(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)
	manager.AddClient(clientSocket)

	manager.DelClient(clientSocket)

	Convey("测试删除客户端", t, func() {
		Convey("长度是否够", func() {
			So(len(manager.ClientIdMap), ShouldEqual, 0)
		})

		Convey("clientId是否存在", func() {
			_, ok := manager.ClientIdMap[clientId]
			So(ok, ShouldBeFalse)
		})
	})
}

func TestCount(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)

	Convey("测试获取客户端数量", t, func() {
		Convey("添加一个客户端后", func() {
			manager.AddClient(clientSocket)
			So(manager.Count(), ShouldEqual, 1)
		})

		Convey("删除一个客户端后", func() {
			manager.DelClient(clientSocket)
			So(manager.Count(), ShouldEqual, 0)
		})

		Convey("再添加两个客户端后", func() {
			manager.AddClient(clientSocket)
			manager.AddClient(clientSocket)
			So(manager.Count(), ShouldEqual, 1)
		})
	})
}

func TestGetByClientId(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)

	Convey("测试通过clientId获取客户端", t, func() {
		Convey("获取一个存在的clientId", func() {
			manager.AddClient(clientSocket)
			_, err := manager.GetByClientId(clientId)
			So(err, ShouldBeNil)
		})

		Convey("获取一个不存在的clientId", func() {
			_, err := manager.GetByClientId("notExistId")
			So(err, ShouldNotBeNil)
		})
	})
}

func TestAddClient2LocalGroup(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)
	manager.AddClient(clientSocket)
	groupName := "testGroup"

	Convey("测试添加分组", t, func() {
		Convey("添加一个客户端到分组", func() {
			manager.AddClient2LocalGroup(&groupName, clientSocket)
			So(len(manager.Groups[groupName]), ShouldEqual, 1)
		})

		Convey("再添加一个客户端到分组", func() {
			manager.AddClient2LocalGroup(&groupName, clientSocket)
			So(len(manager.Groups[groupName]), ShouldEqual, 2)
		})
	})
}

func TestGetGroupClientList(t *testing.T) {
	clientId := "clientId"
	var manager = NewClientManager() // 管理者
	conn := &websocket.Conn{}
	clientSocket := NewClient(clientId, conn)
	manager.AddClient(clientSocket)
	groupName := "testGroup"

	Convey("测试添加分组", t, func() {
		Convey("获取一个存在的分组", func() {
			manager.AddClient2LocalGroup(&groupName, clientSocket)
			clientList := manager.GetGroupClientList(groupName)
			So(len(clientList), ShouldEqual, 1)
		})

		Convey("获取一个不存在的clientId", func() {
			clientList := manager.GetGroupClientList("notExistId")
			So(len(clientList), ShouldEqual, 0)
		})
	})
}
