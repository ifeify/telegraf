//go:generate ../../../tools/readme_config_includer/generator
package azure_eventhub_consumer

import (
	_ "embed"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

type AzureEventHub struct {
	// Configuration
	CheckpointInterval config.Duration `toml:"checkpoint_interval"`
	Log telegraf.Logger `toml:"-"`

	parser telegraf.Parser
}

func (*AzureEventHub) SampleConfig() string {
	return sampleConfig
}

func (e *AzureEventHub) SetParser(parser telegraf.Parser) {
	e.parser = parser
}

func (*AzureEventHub) Init() error {
	return nil
}

func (*AzureEventHub) Gather(telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("azure_eventhub_consumer", func() telegraf.Input {
		return &AzureEventHub{}
	})
}