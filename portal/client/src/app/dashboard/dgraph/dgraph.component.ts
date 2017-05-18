import { Component, OnInit, OnDestroy, Input, ElementRef, ViewChild} from '@angular/core';
import { Graph } from '../../models/graph.model';
import { DashboardService } from '../services/dashboard.service'
import { MenuService } from '../../services/menu.service'
import * as d3 from 'd3';

@Component({
  selector: 'app-dgraph',
  template: `<div class="dgraph" #chart></div>`,
  styleUrls: ['./dgraph.component.css']
})

export class DGraphComponent implements OnInit, OnDestroy {
  @ViewChild('chart') private chartContainer: ElementRef;
  @Input() private graph : Graph;
  private margin: any = { top: 40, bottom: 30, left: 60, right: 20};
  private svg : any
  private x : any;
  private y : any;
  private xAxis: any;
  private yAxis: any;
  private legend : any
  private focus : any
  private element: any
  private created = false

  private chart: any;
  private width: number;
  private height: number;

  constructor(
    private menuService : MenuService,
    private dashboardService : DashboardService) { }

  ngOnInit() {
    this.createGraph();
    this.resizeGraph()
    this.dashboardService.onNewData.subscribe(
      () => {
        //this.svg.remove();
        this.svg.selectAll("*").remove();
        this.updateGraph();
      }
    )
    this.menuService.onWindowResize.subscribe(
      (win) => {
        this.svg.selectAll("*").remove();
        this.resizeGraph()
      }
    );
  }

  ngOnDestroy() {
    this.svg.selectAll("*").remove();
    //this.metricsService.onNewData.unsubscribe()
  }

  createGraph() {
    // set the dimensions and margins of the graph
    this.element = this.chartContainer.nativeElement;
    //console.log("create parent: "+this.element.offsetWidth+","+this.element.offsetHeight)
    //this.width = this.element.offsetWidth - this.margin.left - this.margin.right;
    //this.height = this.element.offsetHeight - this.margin.top - this.margin.bottom;
    this.width = this.graph.width - this.margin.left - this.margin.right;
    this.height = this.graph.height - this.margin.top - this.margin.bottom;
    //console.log("create: "+this.graph.title+": "+this.width+","+this.height)
    this.svg = d3.select(this.element)
      .append('svg')
        //.attr('width', this.element.offsetWidth)
        //.attr('height', this.element.offsetHeight)
        .attr('width',2000)// this.graph.width)
        .attr('height', 2000)//this.graph.height)
      .append("g")
        .attr("transform", "translate(" + this.margin.left + "," + this.margin.top + ")")
    //this.updateGraph()
    this.created=true
  }

  resizeGraph() {
    if (!this.created) {
      return
    }
    this.element = this.chartContainer.nativeElement;
    //console.log("resize parent: "+this.element.offsetWidth+","+this.element.offsetHeight)
    //this.width = this.element.offsetWidth - this.margin.left - this.margin.right;
    //this.height = this.element.offsetHeight - this.margin.top - this.margin.bottom;
    this.width = this.graph.width - this.margin.left - this.margin.right;
    this.height = this.graph.height - this.margin.top - this.margin.bottom;
    console.log("resize: "+this.graph.title+": "+this.width+","+this.height)
    d3.select('svg')
      //.attr('width', this.element.offsetWidth)
      //.attr('height', this.element.offsetHeight)
      .attr('width', this.graph.width)
      .attr('height', this.graph.height)
    //d3.select("g").attr("transform", "translate(" + this.margin.left + "," + this.margin.top + ")")
    //this.updateGraph()
  }

  updateGraph() {
    this.chart = this.svg.append('g')
      .attr('class', 'lines')
      .attr('transform', `translate(${this.margin.left}, ${this.margin.top})`);

    this.x = d3.scaleTime().range([0, this.width]);
    this.y = d3.scaleLinear().range([this.height, 0]);

    // add the X Axis
    if (this.width>80) {
      this.xAxis = this.svg.append("g")
        .attr("class", "axisx")
        .attr("transform", "translate(0," +  this.height + ")")
        .call(d3.axisBottom(this.x).ticks(5));
    }

    // add the Y Axis
    if (this.height>50) {
      this.yAxis = this.svg.append("g")
        .attr("class", "axisy")
        .call(d3.axisLeft(this.y));

      if (this.graph.yTitle != '') {
        this.svg.append("text")
          .attr("class", "y-title")
          .attr("transform", "rotate(-90)")
          .attr("y", 0 - this.margin.left)
          .attr("x", 0 - (this.height / 2))
          .attr("dy", "1em")
          .style("text-anchor", "middle")
          .text(this.graph.yTitle);
        }
    }
  }

}
