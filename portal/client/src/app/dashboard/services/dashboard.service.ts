import { Injectable } from '@angular/core';
import { HttpService } from '../../services/http.service';
import { MenuService } from '../../services/menu.service';
import { Subject } from 'rxjs/Subject'
import { Graph } from '../../models/graph.model';

@Injectable()
export class DashboardService {
    onNewData = new Subject();
    x0 = 20
    y0 = 20
    w0 = 150
    h0 = 100

  constructor(
    private httpService : HttpService,
    private menuService : MenuService) { }

}
