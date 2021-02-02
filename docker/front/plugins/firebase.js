import firebase from "firebase";

if (!firebase.apps.length) {
  const firebaseConfig = {
    apiKey: process.env.apiKey,
    authDomain: `${process.env.projectId}.firebaseapp.com`,
    databaseURL: `https://${process.env.databaseURL}.firebaseio.com`,
    projectId: process.env.projectId,
    storageBucket: `${process.env.projectId}.appspot.com`,
    messagingSenderId: process.env.messagingSenderId,
    appId: process.env.appId,
    measurementId: process.env.measurementId
  };
  firebase.initializeApp(firebaseConfig)
}

const googleProvider = new firebase.auth.GoogleAuthProvider()

export default firebase
export {googleProvider}
