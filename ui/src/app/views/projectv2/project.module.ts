import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from 'app/shared/shared.module';
import { projectV2Routing } from './project.routing';
import { ProjectV2ShowComponent } from 'app/views/projectv2/project.component';
import { ProjectV2TopMenuComponent } from 'app/views/projectv2/top-menu/project.top.menu.component';
import { ProjectV2SidebarComponent } from 'app/views/projectv2/sidebar/sidebar.component';

@NgModule({
    declarations: [
        ProjectV2TopMenuComponent,
        ProjectV2ShowComponent,
        ProjectV2SidebarComponent,
    ],
    imports: [
        SharedModule,
        RouterModule,
        projectV2Routing,
    ],
    schemas: [
        CUSTOM_ELEMENTS_SCHEMA
    ]
})
export class ProjectV2Module {
}
