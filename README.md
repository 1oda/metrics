# Go Demo

### demo4 框架架构
```bash
.
├── common                                  # 通用常量/格式化常量
│   ├── responsestatus.go                   # 错误码
│   └── time.go                             # 时间格式化
├── config                                  # 项目配置
│   └── app.yml                             # 项目配置模块
├── controller                              # 控制层
│   ├── index.go                            # `默认首页`
│   ├── monitor.go                          # 监控项
│   └── process.go                          # 进程项
├── go.mod
├── go.sum
├── log
│   └── app.log
├── main.go                                 # 入口模块
├── model                                   # 模型层
│   ├── core                                # 核心模块
│   │   └── db.go                           # 数据库模块
│   └── service                             # 服务模块
│       ├── monitor.go                      # 监控服务
│       └── process.go                      # 进程服务
├── router                                  # 路由层
│   └── route.go                            # 路由模块
├── static                                  # 静态资源模块
└── utils                                   # 工具/中间件层
    ├── config                              # 配置解析
    │   ├── base.go                         # 基础封装解析模块
    │   └── config.go                       # 全局配置注册模块
    ├── db                                  # 数据库注册模块
    │   └── db.go                           # 数据库注册模块
    ├── log                                 # 日志初始化
    │   └── log.go             、           # 全局日志注册、基础封装日志模块
    ├── response
    │   └── response.go                     # 全局请求响应封装
    └── route             、                # 路由初始化
        └── route.go             、         # 全局路由注册

```

### docker部署MySQL
```bash
docker run --name mysql_3307 \
-v /data/mysql/3307/conf/:/etc/mysql/conf.d \
-v /data/mysql/3307/data:/var/lib/mysql \
-v /data/mysql/3307/log:/root/mysql/logs \
-p 3307:3306 \
--restart always \
-e MYSQL_ROOT_PASSWORD=root \
-itd mysql:5.7


# 登录mysql： 
docker exec -it mysql_3307 bash
mysql -uroot  -P 3306 -proot

# 创建demo4数据库
CREATE DATABASE `demo4` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

# 增加用户和权限
grant all on demo4.* to 'test'@'%' identified by 'test123' ;
```

### docker部署prometheus：
```bash
docker run  -itd --name prometheus \
--restart=always \
-p 9090:9090 \
-v /data/prometheus/conf/prometheus.yml:/etc/prometheus/prometheus.yml  \
-v /data/prometheus/data/:/prometheus \
prom/prometheus
```

### prometheus.yml配置
```bash
global:
  scrape_interval:     5s
  evaluation_interval: 3s

scrape_configs:

  - job_name: prometheus
    static_configs:
      - targets: ['172.23.169.157:9090']
        labels:
          instance: prometheus

  - job_name: server
    static_configs:
      - targets: ['127.0.0.1:8992']
        labels:
          instance: 127.0.0.1


```


### docker部署grafana：
```bash
docker run -itd --name grafana \
--restart=always \
-p 3000:3000 \
-v /data/grafana:/var/lib/grafana \
grafana/grafana
```

### grafana配置json
```json
{
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": {
          "type": "grafana",
          "uid": "-- Grafana --"
        },
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "id": 3,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "CPU使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "continuous-GrYlRd"
          },
          "mappings": [],
          "max": 100,
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 6,
        "w": 7,
        "x": 0,
        "y": 0
      },
      "id": 7,
      "options": {
        "displayMode": "lcd",
        "minVizHeight": 10,
        "minVizWidth": 0,
        "orientation": "horizontal",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "showUnfilled": true
      },
      "pluginVersion": "9.0.4",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_cpu{job=\"server\", name=\"usage\"}",
          "legendFormat": "{{instance}} - {{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "CPU usage",
      "type": "bargauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "CPU使用率",
      "fieldConfig": {
        "defaults": {
          "mappings": [],
          "max": 100,
          "min": 0,
          "thresholds": {
            "mode": "percentage",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "orange",
                "value": 70
              },
              {
                "color": "red",
                "value": 90
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 6,
        "w": 9,
        "x": 7,
        "y": 0
      },
      "id": 3,
      "options": {
        "orientation": "auto",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "text": {
          "valueSize": 28
        }
      },
      "pluginVersion": "9.0.4",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_load{ job=\"server\"}",
          "legendFormat": "{{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Load",
      "type": "gauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "内存使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "continuous-GrYlRd"
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 5,
        "w": 8,
        "x": 16,
        "y": 0
      },
      "id": 10,
      "options": {
        "displayMode": "lcd",
        "minVizHeight": 10,
        "minVizWidth": 0,
        "orientation": "horizontal",
        "reduceOptions": {
          "calcs": [
            "lastNotNull"
          ],
          "fields": "",
          "values": false
        },
        "showUnfilled": true
      },
      "pluginVersion": "9.0.4",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_mem{ job=\"server\", name!=\"usage\"} / 1024",
          "legendFormat": "{{instance}} - {{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "内存（G）",
      "type": "bargauge"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": " BytesRecv &&  bytesSent\n",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 9,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "smooth",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 5,
        "w": 12,
        "x": 0,
        "y": 6
      },
      "id": 5,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "9.0.4",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_network{ job=\"server\"} / 1024 / 1024 / 1024",
          "legendFormat": "{{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "网络",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "CPU使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "percent"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 5,
        "w": 12,
        "x": 12,
        "y": 6
      },
      "id": 8,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_load{job=\"server\"}",
          "legendFormat": "{{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Load",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "CPU使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 6,
        "w": 12,
        "x": 0,
        "y": 11
      },
      "id": 9,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_cpu{job=\"server\", name=\"usage\"}",
          "legendFormat": "{{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "CPU usage",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "内存使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "smooth",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 6,
        "w": 12,
        "x": 12,
        "y": 11
      },
      "id": 11,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_mem{ job=\"server\",name=\"usage\"}",
          "legendFormat": "{{instance}} - {{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "内存使用率 usage",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "内存使用率",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "normal"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 5,
        "w": 12,
        "x": 0,
        "y": 17
      },
      "id": 4,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_mem{job=\"server\"}",
          "legendFormat": "{{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "内存",
      "type": "timeseries"
    },
    {
      "datasource": {
        "type": "prometheus",
        "uid": "gqyN8bgVz"
      },
      "description": "磁盘",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 25,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "smooth",
            "lineWidth": 1,
            "pointSize": 1,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "percent"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 5,
        "w": 12,
        "x": 12,
        "y": 17
      },
      "id": 6,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "pluginVersion": "9.0.4",
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "gqyN8bgVz"
          },
          "editorMode": "code",
          "expr": "prom_server_disk{job=\"server\"} / 1024",
          "legendFormat": " {{name}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "磁盘",
      "type": "timeseries"
    }
  ],
  "schemaVersion": 36,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-5m",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "node",
  "uid": "CEt90fgVz",
  "version": 12,
  "weekStart": ""
}

```
