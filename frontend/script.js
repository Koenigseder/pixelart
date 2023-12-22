class Canvas {
  constructor() {
    this.canvas = document.getElementById("pixelCanvas");
    this.ctx = this.canvas.getContext("2d");

    this.updateCanvas();

    // Listen to tab visibility change
    document.addEventListener("visibilitychange", () => {
      if (document.visibilityState === "visible") {
        this.updateCanvas();
      }
    });
  }

  updateCanvas() {
    this.drawPixels();

    setTimeout(() => {
      // Only update the chart if the tab is active
      if (document.visibilityState === "visible") {
        this.updateCanvas();
      }
    }, 500);
  }

  drawPixels() {
    getPixels().then((res) => {
      const pixels = res.pixels;

      pixels.map((row, i) => {
        row.map((cell, j) => {
          this.ctx.fillStyle = `rgb(${cell[0]} ${cell[1]} ${cell[2]})`;
          this.ctx.fillRect(i, j, 1, 1);
        });
      });
    });
  }
}

new Canvas();
