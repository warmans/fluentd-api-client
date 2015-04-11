package monitoring

import (
	"encoding/json"
	"log"
	"net/http"
)

//Plugins is a collection of Plugin items
type Plugins struct {
	Plugins []Plugin `json:"plugins"`
}

//Plugin represents an active fluentd plugin
type Plugin struct {
	PluginId              string `json:"plugin_id"`
	PluginCategory        string `json:"plugin_category"`
	Type                  string `json:"type"`
	OutputPlugin          bool   `json:"output_plugin"`
	BufferQueueLength     int    `json:"buffer_queue_length"`
	BufferTotalQueuedSize int    `json:"buffer_total_queued_size"`
	RetryCount            int    `json:"retry_count"`
}

//Host represents a td-agent host (address/port)
type Host struct {
	Address   string
	Online    bool
	LastError string
	Plugins   Plugins
}

//Update refreshes the host data including up/down state and plugins
func (h *Host) Update() {

	response, err := http.Get("http://" + h.Address + "/api/plugins.json")

	if err != nil {
		h.handleUpdateError(err)
		return
	}

	plugins := Plugins{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&plugins); err != nil {
		h.handleUpdateError(err)
		return
	}

	h.clearUpdateError()
	h.Plugins = plugins
}

func (h *Host) handleUpdateError(err error) {
	log.Printf("Error querying host %s: %s", h.Address, err.Error())
	h.Online = false
	h.LastError = err.Error()
}

func (h *Host) clearUpdateError() {
	h.Online = true
	h.LastError = ""
}

//NewHost returns a new Host instance. Hosts have Plugins
func NewHost(Address string) *Host {
	return &Host{Address: Address, Online: false, Plugins: Plugins{Plugins: make([]Plugin, 0)}}
}

