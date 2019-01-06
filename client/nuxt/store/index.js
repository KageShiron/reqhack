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
  async createBin(context, params) {
    try{ 
      const { data } = await axios.post(
        `/v1/${params.binname}/create?isPrivate=${params.isPrivate}`
      )
      if (data.success) {
        this.$router.push({ path: `/bin/${name}/` })
      }
    } catch (e) {
      context.commit('updateError', e.response.data.error)
    }
  }
}
