package plantilla

const TemplateContentWithDynamicRoutes = `
import { Route } from '@angular/router';
import { {{.ComponentName | title}}Container } from './components/{{.ComponentName}}-container/{{.ComponentName}}-container';

export const {{.ComponentName }}Routes: Route[] = [
	{
	  path: '',
	  component: {{.ComponentName | title }}Container,
	}
];
`
