import {Component, Input, OnInit} from '@angular/core';
import {ApiService} from '../../api.service';
import {NzTableQueryParams} from 'ng-zorro-antd';
import {ActivatedRoute, Router} from '@angular/router';
import {TabRef} from '../tabs/tabs.component';

@Component({
  selector: 'app-project-strategy',
  templateUrl: './project-strategy.component.html',
  styleUrls: ['./project-strategy.component.scss']
})
export class ProjectStrategyComponent implements OnInit {
  @Input() project: any = {};

  strategies: [];
  total = 0;
  pageIndex = 1;
  pageSize = 10;
  sortField = null;
  sortOrder = null;
  filters = [];
  keyword = '';
  loading = false;

  constructor(private as: ApiService, private router: Router, private routeInfo: ActivatedRoute, private tab: TabRef) {
    tab.name = '自动策略';
  }

  ngOnInit(): void {
    this.tab.name = '项目【' + this.project.name + '】自动策略';
  }

  reload(): void {
    this.pageIndex = 1;
    this.keyword = '';
    this.load();
  }

  load(): void {
    this.loading = true;
    this.as.post('/project/' + this.project.id + '/strategies', {
      offset: (this.pageIndex - 1) * this.pageSize,
      length: this.pageSize,
      sortKey: this.sortField,
      sortOrder: this.sortOrder,
      filters: this.filters,
      keyword: this.keyword,
    }).subscribe(res => {

      this.strategies = res.data;
      this.total = res.total;
    }, error => {
      console.log('error', error);
    }, () => {
      this.loading = false;
    });
  }

  create(): void {
    this.router.navigate(['/admin/project/' + this.project.id + '/strategy/create']);
  }

  edit(c): void {
    this.router.navigate(['/admin/project/' + this.project.id + '/strategy/' + c.id + '/edit']);
  }

  onTableQuery(params: NzTableQueryParams): void {
    const {pageSize, pageIndex, sort, filter} = params;
    this.pageSize = pageSize;
    this.pageIndex = pageIndex;
    const currentSort = sort.find(item => item.value !== null);
    this.sortField = (currentSort && currentSort.key) || null;
    this.sortOrder = (currentSort && currentSort.value) || null;
    this.filters = filter;
    this.load();
  }

  search(): void {
    this.pageIndex = 1;
    this.load();
  }
}
