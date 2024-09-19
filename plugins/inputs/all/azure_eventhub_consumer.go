//go:build !custom || inputs || inputs.azure_eventhub_consumer

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/azure_eventhub_consumer" // register plugin
