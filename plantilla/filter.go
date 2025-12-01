package plantilla

const TemplateContenttWithDynamicFilterComponent = `
import { Component, EventEmitter, Output } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-{{.ComponentName}}-filter',
  imports: [ReactiveFormsModule],
  templateUrl: './{{.ComponentName}}-filter.html',
  styleUrl: './{{.ComponentName}}-filter.css',
})

export class {{.ComponentName | title }}Filter {
	form = new FormGroup({
   		search: new FormControl('', []),
    })

    @Output() paramsChange = new EventEmitter<{ action: string, values: any }>();

    onFilterChange() {}

    onSeach(){}

    onAddChange() {}
}
`

const TemplateContenttWithDynamicFilterHtml = `
<form class="flex justify-between" [formGroup]="form">
  <div class="w-70">
    <div class="form-control-group icon-right">
      <i
        class="bx bx-search-alt-2 text-gray-400 text-2xl! cursor-pointer!"
        tabindex="0"
        role="button"
        (click)="onSeach()"
      ></i>
      <input
        type="text"
        id="name"
        placeholder="Buscar por nombre"
        class="form-control form-control-md"
        formControlName="search"
        autocomplete="off"
      />
    </div>
  </div>
</form>
`

const TemplateContenttWithDynamicFilterCss = ``
