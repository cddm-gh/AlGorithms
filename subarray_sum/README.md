<div class="col col-12 col-sm-10 text-justify">
      <h1>0x25 - Suma de un subarreglo grande</h1>
      <div>
        <p>Supongamos que tenemos un arreglo llamado <code>A</code>.</p>
<p>Un subarreglo de <code>A</code> está definido como cualquier bloque de elementos consecutivos en <code>A</code>. Por ejemplo, si <span class="nowrap"><code>A = [5, -2, 1, 4]</code></span>, entonces <code>A</code> tiene 10 subarreglos que son:</p>
<ul>
<li><code>[5]</code></li>
<li><code>[5, -2]</code></li>
<li><code>[5, -2, 1]</code></li>
<li><code>[5, -2, 1, 4]</code></li>
<li><code>[-2]</code></li>
<li><code>[-2, 1]</code></li>
<li><code>[-2, 1, 4]</code></li>
<li><code>[1]</code></li>
<li><code>[1, 4]</code></li>
<li><code>[4]</code>.</li>
</ul>
<p>Por otro lado, <span class="nowrap"><code>[5, -2, 4]</code></span> <strong>no</strong> es un subarreglo de <code>A</code> porque estos elementos no aparecen de forma <strong>consecutiva</strong> en <code>A</code>.</p>
<p>Una manera de describir un subarreglo es usando dos índices de <code>A</code> que describen dónde empieza y dónde termina el subarreglo. Llamemos <code>p</code> al índice izquierdo y <code>q</code> al índice derecho. Por ejemplo:</p>
<ul>
<li>El subarreglo <span class="nowrap"><code>[5, -2]</code></span> se describe con <span class="nowrap"><code>p = 0</code></span> y <span class="nowrap"><code>q = 1</code></span>;</li>
<li>El subarreglo <span class="nowrap"><code>[-2, 1, 4]</code></span> se describe con <span class="nowrap"><code>p = 1</code></span> y <span class="nowrap"><code>q = 3</code></span>;</li>
<li>El subarreglo <code>[1]</code> se describe con <span class="nowrap"><code>p = 2</code></span> y <span class="nowrap"><code>q = 2</code></span>;</li>
<li>El subarreglo <code>[5, -2, 1, 4]</code> se describe con <span class="nowrap"><code>p = 0</code></span> y <span class="nowrap"><code>q = 3</code></span>.</li>
</ul>
<p>Escribe un programa que recibe varias "consultas" de la forma <span class="nowrap"><code>p, q</code></span> y para cada consulta calcula la suma de todos los elementos del subarreglo que empieza en <code>p</code> y termina en <code>q</code>.</p>
<h2>Entrada</h2>
<p>La entrada contiene varias líneas:</p>
<ul>
<li>La primera línea contiene <code>N</code>, el número de elementos en el arreglo.</li>
<li>La segunda línea contiene <code>N</code> enteros <code>A<sub>i</sub></code> para <span style="white-space: nowrap">0 ≤ <code>i</code> &lt; <code>N</code></span>, los elementos del arreglo. Estos números están separados por espacios. Está garantizado que <code>A<sub>i</sub></code> satisface <span style="white-space: nowrap">-10<sup>4</sup> ≤ <code>A<sub>i</sub></code> ≤ 10<sup>4</sup></span>.</li>
<li>La tercera línea contiene <code>C</code>, el número de consultas.</li>
<li>Luego vienen <code>C</code> líneas, cada una con un consulta descrita por dos índices <code>p</code> y <code>q</code> separados por un espacio. Está garantizado que <code>0 ≤ p ≤ q &lt; N</code> (es decir, <code>p</code> y <code>q</code> son índices válidos).</li>
</ul>
<h2>Salida</h2>
<p>La salida debe tener exactamente <code>C</code> líneas (una por consulta).</p>
<p>Para cada consulta, escribe una línea con la suma de los elementos del subarreglo que empieza en <code>p</code> y termina en <code>q</code>.</p>
<h2>Entrada de ejemplo</h2>
<pre><code>4
5 -2 1 4
4
0 1
1 3
2 2
0 3
</code></pre>
<h2>Salida de ejemplo</h2>
<pre><code>3
3
1
8
</code></pre>
<h2>Explicación de la entrada y salida de ejemplo</h2>
<p>La entrada nos da un arreglo <span class="nowrap"><code>A = [5, -2, 1, 4]</code></span> y 4 consultas:</p>
<ul>
<li>La primera consulta (<code>0 1</code>) corresponde al subarreglo <span class="nowrap"><code>[5 -2]</code></span>. La suma de estos elementos es 3.</li>
<li>La segunda consulta (<code>1 3</code>) corresponde al subarreglo <span class="nowrap"><code>[-2 1 4]</code></span>. La suma de estos elementos también es 3 (por casualidad).</li>
<li>La tercera consulta (<code>2 2</code>) corresponde al subarreglo <code>[1]</code>. La suma de este único elemento es 1.</li>
<li>La cuarta y última consulta (<code>0 3</code>) corresponde al arreglo completo, <span class="nowrap"><code>[5 -2 1 4]</code></span>. La suma de estos elementos es 8.</li>
</ul>
<h2>Restricciones</h2>
<ul>
<li>En 50% de los casos, <code>1 ≤ N ≤ 100</code> y <code>1 ≤ C ≤ 100</code>.</li>
<li>En el otro 50% de los casos, <span class="nowrap"><code>1 ≤ N ≤ 50,000</code></span> y <span class="nowrap"><code>1 ≤ C ≤ 50,000</code></span>.</li>
