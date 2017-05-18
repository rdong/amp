import { Component, OnInit } from '@angular/core';
import { MenuService } from '../services/menu.service';
import { ActivatedRoute } from '@angular/router';
import { AppWindow } from '../models/app-window.model';
import { Graph } from '../models/graph.model';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})

export class DashboardComponent implements OnInit {
  graphs : Graph[] = []
  dashboardName = "default"
  periodRefreshLabel = "30 seconds"
  graphPanelHeight = 250
  graphPanelWidth = 500

  constructor(
    private menuService : MenuService,
    //public metricsService : DashboardComponentService,
    private route: ActivatedRoute) {
  }

  ngOnInit() {
    this.menuService.setItemMenu('dashboard', 'View')
    this.resizeGraphs(this.menuService.appWindow)
    this.menuService.onWindowResize.subscribe(
      (win) => {
        this.resizeGraphs(win)
      }
    )
  }

  resizeGraphs(win : AppWindow) {
    //let cww = $('.graph-container').width()
    let cww = win.width-25-this.menuService.paddingLeftMenu
    let chh = win.height- 200;
    //console.log("Window: "+win.width+","+win.height)
    //console.log("Container: "+cww+","+chh)
    this.graphPanelHeight = chh
    this.graphPanelWidth = cww
  }

  setRefreshPeriod(period : string, refrsh : string) {

  }

  addGraph(type : string) {
    this.graphs.push(new Graph(20, 20, 100, 100, type, [''], '',''))
  }

}
