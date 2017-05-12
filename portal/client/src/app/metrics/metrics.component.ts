import { Component, Input, OnInit, OnDestroy } from '@angular/core';
import { MenuService } from '../services/menu.service';
import { MetricsService } from './services/metrics.service';
import { DockerStacksService } from '../docker-stacks/services/docker-stacks.service';
import { StatsRequest } from './models/stats-request.model';
import { ActivatedRoute } from '@angular/router';
import { AppWindow } from '../models/app-window.model';
import { Graph } from './models/graph.model';
declare var $: any

@Component({
  selector: 'app-dashboard',
  templateUrl: './metrics.component.html',
  styleUrls: ['./metrics.component.css']
})
export class MetricsComponent implements OnInit, OnDestroy {
  periodTimeLabel = "10 min"
  periodRefreshLabel = "30 sec"
  routeSub : any
  name = ""
  dashboardName = ""
  refName =""
  graphPanelHeight = 250
  graphs : Graph[] = []

  constructor(
    private menuService : MenuService,
    public metricsService : MetricsService,
    private route: ActivatedRoute,
    private dockerStacksService : DockerStacksService) {
      this.graphs.push(new Graph(0, 0, 0, 0, ['cpu-usage'], "cpu", "usage %"))
      this.graphs.push(new Graph(0, 0, 0, 0, ['mem-usage'], "memory", "usage MB"))
      this.graphs.push(new Graph(0, 0, 0, 0, ['net-total-bytes'], "network", "total bytes"))
      this.graphs.push(new Graph(0, 0, 0, 0, ['io-total'], "disk io", "total bytes"))
      this.resizeGraphs(this.menuService.appWindow)
    }

  ngOnInit() {
    this.resizeGraphs(this.menuService.appWindow)
    this.menuService.onWindowResize.subscribe(
      (win) => {
        this.resizeGraphs(win)
      }
    )
    this.menuService.setItemMenu('metrics', 'View')
    this.routeSub = this.route.params.subscribe(params => {
      let object = params['object']//'stack' or 'service' or 'task'
      let type = params['type'];//'single' or 'multi'
      let ref = params['ref'];//stackName or serviceId or taskId
      //console.log("object="+this.object+" type="+this.type+" ref="+this.ref)
      this.metricsService.set(object, type, ref)
      this.computeNames()

      let req = new StatsRequest()
      req.stats_cpu = true
      req.stats_mem = true
      req.stats_io = true
      req.stats_net = true
      req.period = this.metricsService.timePeriod
      req.time_group = this.metricsService.timeGroup
      if (object=='stack') {
        if (type == 'single') {
          req.filter_stack_name = ref
          this.menuService.setItemMenu('metrics', 'View stack')
        }
        if (type == 'multi') {
          req.group = 'stack_name'
          this.menuService.setItemMenu('metrics', 'View stacks')
        }
      }
      if (object=='service') {
        if (type == 'single') {
          req.filter_service_name = ref
          this.menuService.setItemMenu('metrics', 'View service')
        }
        if (type == 'multi') {
          req.group = 'service_name'
          req.filter_stack_name = ref
          this.menuService.setItemMenu('metrics', 'View services')
        }
      }
      if (object=='task') {
        if (type == 'single') {
          req.filter_task_id = ref
          this.menuService.setItemMenu('metrics', 'View container')
        }
        if (type == 'multi') {
          req.group = 'container_short_name'
          req.filter_service_name = ref
          this.menuService.setItemMenu('metrics', 'View containers')
        }
      }
      //console.log(req)
      this.metricsService.setHistoricRequest(req, this.metricsService.periodRefresh)
    })
  }

  ngOnDestroy() {
    this.metricsService.cancelRequests()
  }

  setTimePeriod(label : string, period, group : string) {
    this.periodTimeLabel = label
    this.metricsService.setTimePeriod(period, group)
  }

  setRefreshPeriod(label : string, period : string) {
    this.periodRefreshLabel = label
    this.metricsService.setRefreshPeriod(period)
  }

  returnBack() {
    this.menuService.returnToPreviousPath()
  }

  computeNames() {
    //console.log("compute.name: "+[this.metricsService.object, this.metricsService.type, this.metricsService.ref])
    this.dashboardName=this.metricsService.object
    this.refName = this.metricsService.ref
    if (this.metricsService.object == 'global') {
      this.dashboardName="Global"
      this.refName = ""
    }
    if (this.metricsService.object == 'stack') {
      this.dashboardName="Stacks"
    }
    if (this.metricsService.object == 'service') {
      this.dashboardName="Services of stack:"
    }
    if (this.metricsService.object == 'task') {
      this.dashboardName="Containers of service:"
    }
  }

  //[style.height.px]="parentdiv.offsetHeight"
  resizeGraphs(win : AppWindow) {
    let cww = $('.graph-container').width()
    let chh = win.height- 240;
    //console.log("Window: "+win.width+","+win.height)
    //console.log("Container: "+cww+","+chh)
    this.graphPanelHeight = chh
    let xx=10
    let yy=10
    for (let graph of this.graphs) {
      graph.width = Math.floor(cww/2 - 20)
      graph.height = Math.floor(chh/2 - 15)
      graph.x = Math.floor(xx)
      graph.y = Math.floor(yy)
      xx = xx + cww/2 - 10
      if (xx + graph.width > cww) {
        xx = 10
        yy = yy + chh/2 - 15
      }
      //console.log(graph)
    }
  }

}
