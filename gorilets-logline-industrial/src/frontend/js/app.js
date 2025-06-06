import { connectWs } from './ws.js';
import { renderLoglines } from './render.js';

fetch('ui/ui_flip.logline')
  .then(res => res.text())
  .then(text => {
    document.getElementById('app').innerText = 'FlipApp carregado com spans:\n' + text;
  });

connectWs();
renderLoglines();
