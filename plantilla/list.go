package plantilla

const TemplateContenttWithDynamicListComponent = `
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { BadgeModule } from 'primeng/badge';
import { TableModule } from 'primeng/table';
import { UiButtonDropdown, UiPagination } from '@org/ui-components';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-{{.ComponentName}}-list',
  imports: [TableModule, CommonModule, BadgeModule, UiPagination, UiButtonDropdown],
  templateUrl: './{{.ComponentName}}-list.html',
  styleUrl: './{{.ComponentName}}-list.css',
})

export class {{.ComponentName | title }}List {
  @Input() items: any[] = [];

  @Output() pageChange = new EventEmitter<number>();
  @Output() editChange = new EventEmitter<any>();
  @Output() deleteChange = new EventEmitter<any>();

  onPageChange(page: number) {
    this.pageChange.emit(page);
  }
  onEdit(item: any) {
    this.editChange.emit(item);
  }
  onDelete(item: any) {
    this.deleteChange.emit(item);
  }
}
`

const TemplateContenttWithDynamicListHtml = `
<div class="flex justify-center">
  <div class="w-full">
  </div>
</div>
`

const TemplateContenttWithDynamicListCss = ``
