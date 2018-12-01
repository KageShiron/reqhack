import axios from 'axios'

export const state = () => ({
  binname: Math.random()
    .toString(36)
    .slice(-8),
  error: ''
})

export const mutations = {
  updateBinName(state, value) {
    state.binname = value
  },
  updateError(state, value) {
    state.error = value
  }
}

export const actions = {
  async createBin(context, name) {
    const { data } = await axios.post(`/v1/${name}/create`)
    if (data.success) {
      this.$router.push({ path: `/bin/${name}/` })
    } else {
      context.commit('updateError', data)
    }
  }
}
