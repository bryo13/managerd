<script>
  import { Test1 } from "../wailsjs/go/store/DB";
  
  let worked
  Test1().then((sucess) => {
     worked = sucess;
  });
</script>

<main>
  <p>{worked}</p>
</main>
