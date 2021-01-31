import firebase from "firebase";

if (!firebase.apps.length) {
  const firebaseConfig = {
  };
  firebase.initializeApp(firebaseConfig)
}

const googleProvider = new firebase.auth.GoogleAuthProvider()

export default firebase
export {googleProvider}
