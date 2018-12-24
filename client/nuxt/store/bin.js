import axios from 'axios'

export const state = () => ({
  items: {},
  headerActiveTab: 'table'
})

export const mutations = {
  update_bin(state, value) {
    state.items[value.name] = value.data
  },
  update_headerActiveTab(state, value) {
    state.headerActiveTab = value
  }
}

export const actions = {
  async fetch_bin(context, name) {
    const { data } = await axios.get(`/v1/${name}/items`)
    context.commit('update_bin', { name, data: data.reverse() })
  }
}
