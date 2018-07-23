package juniper

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestParse(t *testing.T) {
	config := `
	node0 {
        system {
            host-name Cloud_BSA_3A01A_DMZ3400_01;
            backup-router 30.2.100.200 destination 0.0.0.0/0;
            tacplus-server {
                x.x.x.1 source-address x.x.x.x;
                x.x.x.2 source-address x.x.x.x;
            }
            accounting {
                destination {
                    tacplus {
                        server {
                            x.x.x.1 source-address x.x.x.x;
                            x.x.x.2 source-address x.x.x.x;
                        }
                    }
                }
            }
            syslog {
                source-address x.x.x.x;
            }
        }
        interfaces {
            fxp0 {
                unit 0 {
                    family inet {
                        address x.x.x.x/22;
                    }
                }
            }
        }
    }
`

	config1 := `
    alg {
        dns disable;
        h323 disable;
        msrpc disable;
        sunrpc disable;
        rsh disable;
        rtsp disable;
        sql disable;
        talk disable;
        tftp disable;
        pptp disable;
    }
`

	config2 := `
	rule-set Cloud_To_CT {
		from zone Internet;
		to zone CT;
		rule DNS {
			match {
				source-address [ x.x.x.x/32 x.x.x.x/32 ];
			}
			then {
				source-nat {
					pool {
						srcnat_pool_103_28_214_254;
					}
				}
			}
		}
	}
`
	config3 := `
	PUBLIC {
        instance-type virtual-router;
        interface reth1.104;
        interface reth1.105;
        interface reth1.107;
	}
`
	//bs, _ := ioutil.ReadFile("/path/to/juniper.conf")
	//config := string(bs)
	for _, c := range []string{config, config1, config2, config3} {
		r, _ := Parse(c)
		j, _ := json.Marshal(r)
		fmt.Println(string(j))
	}

}
