import { Component, OnInit } from '@angular/core';
import {ApiService} from '../../api.service';
import {VariableEditComponent} from '../variable-edit/variable-edit.component';
import {NzDrawerService, NzTableQueryParams} from 'ng-zorro-antd';

@Component({
  selector: 'app-variable',
  templateUrl: './variable.component.html',
  styleUrls: ['./variable.component.scss']
})
export class VariableComponent implements OnInit {

  variables: [];
  total = 0;
  pageIndex = 1;
  pageSize = 10;
  sortField = null;
  sortOrder = null;
  filters = [];
  keyword = '';
  loading = false;

  statusFilters = [{text: '启动', value: 1}];


  constructor(private as: ApiService, private drawer: NzDrawerService) {
  }

  ngOnInit(): void {
  }

  reload(): void {
    this.pageIndex = 1;
    this.keyword = '';
    this.load();
  }

  load(): void {
    this.loading = true;
    this.as.post('variables', {
      offset: (this.pageIndex - 1) * this.pageSize,
      length: this.pageSize,
      sortKey: this.sortField,
      sortOrder: this.sortOrder,
      filters: this.filters,
      keyword: this.keyword,
    }).subscribe(res => {

      this.variables = res.data;
      this.total = res.total;
    }, error => {
      console.log('error', error);
    }, () => {
      this.loading = false;
    });
  }

  edit(c?): void {
    this.drawer.create({
      nzTitle: c ? '编辑' : '创建',
      nzMaskClosable: false,
      nzWidth: 500,
      nzContent: VariableEditComponent,
      nzContentParams: {
        variable: c || {}
      }
    });
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
