package plantilla

const TemplateContenttWithDynamicContainerComponent = `
import { Component, OnInit, inject } from '@angular/core';
import { handleResponsePaginableError, NotificationUtilService, Paginable } from '@org/utils';
import { ToastModule } from 'primeng/toast';
import { CardModule } from 'primeng/card';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { MessageService } from 'primeng/api';
import {
  BreadcrumbItem,
  UiBreadcrumbComponent,
  UiConfirmDialog,
} from '@org/ui-components';
import { {{.ComponentName | title }}Filter } from '../{{.ComponentName}}-filter/{{.ComponentName}}-filter';
import { {{.ComponentName | title }}List } from '../{{.ComponentName}}-list/{{.ComponentName}}-list';


@Component({
  selector: 'app-{{.ComponentName}}-container',
  imports:[
    UiBreadcrumbComponent,
  	CardModule,
    ToastModule,
    UiConfirmDialog,
    {{.ComponentName | title }}Filter,
    {{.ComponentName | title }}List,
  ],
  providers: [NotificationUtilService],
  templateUrl: './{{.ComponentName}}-container.html',
  styleUrls: ['./{{.ComponentName}}-container.css']
})
export class {{.ComponentName | title }}Container implements OnInit {
 items: BreadcrumbItem[] = [];
 ref!: DynamicDialogRef | null;
 dialogService = inject(DialogService);
 messageService = inject(MessageService);
 notificationUtilService = inject(NotificationUtilService);

  constructor() { }

  ngOnInit(): void {
  }

}
`

const TemplateContenttWithDynamicContainerHtml = `
<div>
  <lib-ui-breadcrumb [items]="items"></lib-ui-breadcrumb>
  <p-card class="!card-no-padding-top bg-white!">
    <p class="text-2xl font-semibold flex items-center">
        <i class="bx bx-list-ul text-green-600 mr-2 leading-none translate-y-0.5"
        ></i>
    </p>
    <br />
    <app-{{.ComponentName}}-filter
      (paramsChange)="changeParams($event)"
    ></app-{{.ComponentName}}-filter>
    <br />
    <app-{{.ComponentName}}-list
      (editChange)="onEdit($event)"
      (deleteChange)="onDelete($event)"
    ></app-{{.ComponentName}}-list>
  </p-card>
</div>
`

const TemplateContenttWithDynamicContainerCss = ``
