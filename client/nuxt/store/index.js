import axios from 'axios'

export const state = () => ({
  error: null
})

export const mutations = {
  updateError(state, value) {
    state.error = value
  }
}

export const actions = {
  async createBin(context, params) {
    axios
      .post(`/v1/${params.binname}/create?isPrivate=${params.isPrivate}`)
      .then(res => {
        if (res.data.success) {
          if (res.data.secret) {
            context.commit('bin/update_bin', {
              name: params.binname,
              data: { secret: res.data.secret }
            })
          }
          this.$router.push({
            path: `/bin/${params.binname}${
              res.data.secret ? '?secret=' + res.data.secret : ''
            }`
          })
        }
      })
      .catch(e => {
        console.log(e)
        context.commit('updateError', e.response.data.error)
      })
  }
}
