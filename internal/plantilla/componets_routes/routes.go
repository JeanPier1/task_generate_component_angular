package componets_routes

const TemplateContentWithDynamicRoutes = `
import { Route } from '@angular/router';
import { {{.ComponentName | title}}Container } from './components/{{.ComponentName}}-container/{{.ComponentName}}-container';

export const {{.ComponentName | title }}Routes: Route[] = [
	{
	  path: '',
	  component: {{.ComponentName | title }}Container,
	}
];
`
