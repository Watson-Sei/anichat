export default function ({ store, redirect }) {
  if(!store.getters['isAdminIn']) {
    return redirect("/")
  }
}
