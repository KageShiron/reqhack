import axios from 'axios'

export const state = () => ({
  binname: Math.random()
    .toString(36)
    .slice(-8),
  created: ''
})

export const mutations = {
  updateBinName(state, value) {
    state.binname = value
  },
  updateCreated(state, value) {
    state.created = value
  }
}

export const actions = {
  async createBin(context, name) {
    const { data } = await axios.post(`/v1/${name}/create`)
    if (data.success) {
      context.commit('updateCreated', name)
    } else {
      context.commit('updateCreated', 'error')
    }
  }
}
