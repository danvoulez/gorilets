fetch("ui/flipapp_ui.logline")
  .then(res => res.text())
  .then(text => {
    document.getElementById("app").innerText = "FlipApp carregado com spans:\n" + text;
  });
