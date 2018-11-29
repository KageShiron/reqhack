import { action, observable, runInAction } from "mobx";
import { Item } from "../features/bin/types";

export default class BinStore {
  @observable public name: string;
  @observable public itemList: Item[] = [];

  @action
  public fetchItems() {
    fetch(`/v1/${this.name}/items`)
      .then(res => res.json())
      .then(data => {
        runInAction(() => {
          this.itemList = data;
        });
      });
  }
}

export const binStore = new BinStore();
