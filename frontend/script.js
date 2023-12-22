class Canvas {
  constructor() {
    this.canvas = document.getElementById("pixelCanvas");
    this.ctx = this.canvas.getContext("2d");
    this.socket = new WebSocket("ws://localhost:8080/api/ws");

    this.drawCanvas();
  }

  drawCanvas() {
    this.socket.onmessage = (msg) => {
      const pixels = JSON.parse(msg.data).pixels;

      pixels.map((row, i) => {
        row.map((cell, j) => {
          this.ctx.fillStyle = `rgb(${cell[0]} ${cell[1]} ${cell[2]})`;
          this.ctx.fillRect(i, j, 1, 1);
        });
      });
    };
  }
}

new Canvas();
