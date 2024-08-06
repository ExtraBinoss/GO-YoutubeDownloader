import './style.css'
import('./appstyles.css').then(() => {
  console.log('CSS dynamically loaded');
});
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app')
})

export default app
