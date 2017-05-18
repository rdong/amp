import { Input, HostListener, Directive, HostBinding } from '@angular/core';


@Directive({
  selector: '[appDraggable]'
})
export class MovableDirective {
  constructor() {

  }

  @HostBinding('movable')
  get movable() {
    return true;
  }

  @Input()
  /*
  set appDraggable(options: DraggablOptions) {
    if (options) {
      this.options = options;
    }
  }
  */

  //private options: DraggableOptions = {};

  @HostListener('dragstart', ['$event'])
  onDragStart(event) {
    //const { zone = 'zone', data = {} } = this.options;

    //this.dragService.startDrag(zone);

    //event.dataTransfer.setData('Text', JSON.stringify(data));
  }
}
