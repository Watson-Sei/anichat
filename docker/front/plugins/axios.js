import axios from "axios";

export default function ({ $axios, redirect }) {
  $axios.onError(error => {
    if (error.response.status === 401 && error.response.data.message === "Token has expired") {
      const params = new URLSearchParams();
      params.append("grant_type", "refresh_token")
      params.append("refresh_token", localStorage.getItem('refresh_token'))
      axios.post(`https://securetoken.googleapis.com/v1/token?key=${process.env.apiKey}`, params, {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        }
      }).then((response) => {
        localStorage.removeItem('access_token')
        localStorage.setItem('access_token', response.data.access_token)
        localStorage.removeItem('refresh_token')
        localStorage.setItem('refresh_token', response.data.refresh_token)
        axios(error.config)
      }).catch(() => {
        redirect('/room')
      })
    }
  })
}
