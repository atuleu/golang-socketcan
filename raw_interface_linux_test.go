//go:build linux

package socketcan

import (
	"os/exec"
	"slices"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RawInterfaceLinuxSuite struct{}

var _ = Suite(&RawInterfaceLinuxSuite{})

var VCanInterfaceNames []string = []string{
	"vcan_test0",
	"vcan_test1",
}

func (s *RawInterfaceLinuxSuite) SetUpSuite(c *C) {
	for _, name := range VCanInterfaceNames {
		cmd := exec.Command("sudo", "ip", "link", "add", "dev", name, "type", "vcan")
		out, err := cmd.CombinedOutput()
		c.Assert(err, IsNil, Commentf("creating interface %s: %s", name, string(out)))
	}
}

func (s *RawInterfaceLinuxSuite) TearDownSuite(c *C) {
	for _, name := range VCanInterfaceNames {
		cmd := exec.Command("sudo", "ip", "link", "delete", name)
		out, err := cmd.CombinedOutput()
		c.Assert(err, IsNil, Commentf("deleting interface %s: %s", name, string(out)))
	}

}

func (s *RawInterfaceLinuxSuite) TestList(c *C) {
	availables, err := ListCANInterfaces()
	c.Assert(err, IsNil)
	c.Check(len(availables) >= len(VCanInterfaceNames), Equals, true,
		Commentf("Expected len(availables) >= %d, got %d", len(VCanInterfaceNames), len(availables)))

	for _, name := range VCanInterfaceNames {
		c.Check(slices.Contains(availables, name), Equals, true, Commentf("availables (%s) should contain %s", availables, name))
	}
}
