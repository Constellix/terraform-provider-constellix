locals {
  #----------------------------------------------------------------------------
  # Workspace Defaults
  #----------------------------------------------------------------------------
  note    = "Test created with terraform"
  domains = ["constellix-terraform-test.com"]

  #----------------------------------------------------------------------------
  # Records
  #----------------------------------------------------------------------------

  records = {
    a = {
      www = ["1.2.3.4"]
    }
    aaaa = {
      www = ["5:0:0:0:0:0:1:6"]
    }
    aname = {
      root = ["google.com."]
    }
    cname = {
      google = ["google.com"] #only put 1 value in this
    }
    hinfo = {
      ###### cpu,os
      _tcp = ["quad core,linux2"]
    }
    httpredirection = {
      red = ["https://www.google.com"] #only put 1 value in this
    }
    mx = {
      ###### level,value
      root = ["10,mail.example.com"]
    }
    naptr = {
      ###### order,preference,flags,service,regular_expression,replacement
      root = ["10,100,s,SIP+D2U,hello,foobar.example.com."]
    }
    caa = {
      # 1 for [ Custom ], 2 for [ No Provider ], 3 for Comodo, 4 for Digicert, 5 for Entrust, 6 for GeoTrust, 7 for Izenpe, 8 for Lets Encrypt, 9 for Symantec, 10 for Thawte
      ###### caa_provider_id,tag,data,flag
      root = ["3,issue,como.com,0"]
    }
    cert = {
      ###### certificate_type,key_tag,certificate,algorithm
      root = ["20,30,certificate1,100"]
    }
    ptr = {
      root = ["10"]
    }
    rp = {
      ###### mailbox,txt
      root = ["one.com,domain.com"]
    }
    spf = {
      root = ["This is depreciated"]
    }
    srv = {
      ###### value,port,priority,weight
      _tcp = ["www.google.com,8888,65,20"]
    }
    txt = {
      root = ["v=spf1"]
    }
    ns = {
      ns = ["ns0.dnsmadeeasy.com."]
    }
  }

  #----------------------------------------------------------------------------
  # Pools
  #----------------------------------------------------------------------------
  pools = {
    a = {
      test_a_pool1 = {
        values = [
          {
            value        = "142.250.191.110"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          },
          {
            value        = "142.250.189.206"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          }
        ]
      }
    }
    cname = {
      test_cname_pool1 = {
        values = [
          {
            value        = "google.com"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          },
          {
            value        = "google.com"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          }
        ]
      }
    }
    aaaa = {
      test_aaaa_pool1 = {
        values = [
          {
            value        = "2607:f8b0:4005:810:0:0:0:200e"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          },
          {
            value        = "2607:f8b0:4009:814::200e"
            weight       = 20
            policy       = "followsonar"
            disable_flag = false
            fqdn         = "google.com"
          }
        ]
      }
    }
  }
}