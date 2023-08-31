package manager
// SPDX-License-Identifier: LGPL-2.1-or-later

import (
        "fmt"
        "os"
        "log"
        "github.com/godbus/dbus"
)

// Instance type refers to the ConfigMap object
type Instance struct {
	Conn	   string
	Interface  string
	BusObject  string
}

func (* Instance) Connect() {
        conn, err := dbus.ConnectSystemBus()
        if err != nil {
                fmt.Fprintln(os.Stderr, "Failed to connect to d-bus:", err)
                os.Exit(1)
        }

	Instance.Conn = conn
        Instance.BusObject = conn.Object(
		common.BC_DBUS_INTERFACE,
		common.BC_OBJECT_PATH
	)
        //defer conn.Close()
}

func (* Instance) ListNodes() {
	METHOD_LISTNODES := "org.eclipse.bluechi.Manager.ListNodes"

        var nodes [][]interface{}

        err = busObject.Call(METHOD_LISTNODES, 0).Store(&nodes)
        if err != nil {
		fmt.Fprintln(os.Stderr, "unable to list nodes: ", err)
                os.Exit(1)
        }

	for _, node := range nodes {
		fmt.Println(node)
	}
}

//conn.BusObject().Call("org.eclipse.bluechi.Manager.ListUnits", 0).Store(&s)
//conn.BusObject().Call("org.eclipse.bluechi.Manager.GetNode", 0).Store(&s)
