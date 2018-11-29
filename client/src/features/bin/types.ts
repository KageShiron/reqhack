import { observable } from "mobx";

export class Item {
  @observable public time: string;
  @observable public method: string;
  @observable public protocol: string;
  @observable public header: Array<{ [key: string]: string }>;
  @observable public body: string;
  @observable public host: string;
  @observable public postform: any;
  @observable public multipartform: any;
  @observable public remoteaddr: string;
  @observable public requesturi: string;
}
