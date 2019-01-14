import axios from 'axios'

export const state = () => ({
  items: {},
  headerActiveTab: 'table'
})

export const mutations = {
  update_bin(state, value) {
    if(!value.name) throw "Bad update bin";
    state.items[value.name] = Object.assign(value.data,state.items[value.name])
  },
  update_headerActiveTab(state, value) {
    state.headerActiveTab = value
  }
}

export const actions = {
  async fetch_bin(context, params) {
    const secret = params.secret || ( context.state.items[params.name] && context.state.items[params.name].secret) || ""
    const { data } = await axios.get(
      `/v1/${params.name}/items?secret=${secret}`
    )
    context.commit('update_bin', {
      name: params.name,
      data: { data: data.reverse() }
    })
  }
}
