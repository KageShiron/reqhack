import axios from 'axios'

export const state = () => ({
  binname: Math.random()
    .toString(36)
    .slice(-8),
  error: null
})

export const mutations = {
  updateError(state, value) {
    state.error = value
  }
}

export const actions = {
  async createBin(context, name) {
    try {
      const { data } = await axios.post(`/v1/${name}/create`)
      this.$router.push({ path: `/bin/${name}/` })
    } catch (e) {
      context.commit('updateError', e.response.data.error)
    }
  }
}
