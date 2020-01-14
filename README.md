# PegNet Consensus Math Simulator

A small utility to simulate pegnet consensus voting by simulating 200,000 voters, half of which vote "yay". The output is 100 blocks of sequential moving averages, printed in comma separated rows.

## Sample Output

* Target: The real ratio that the other values try to estimate
* Block: The ratio inside of the single block that was added
* SMA-#: The simple moving average for a window of that many blocks

```
Target,Block,SMA-1,SMA-6,SMA-12,SMA-18,SMA-24,SMA-30,SMA-36,SMA-42,SMA-48,SMA-54,SMA-60,SMA-66,SMA-72,SMA-78,SMA-84,SMA-90,SMA-96,SMA-102,SMA-108,SMA-114,SMA-120,SMA-126,SMA-132,SMA-138,SMA-144
0.5,0.66,0.66,0.55,0.5083,0.5122,0.5033,0.5,0.5,0.4976,0.505,0.5015,0.4993,0.4973,0.5,0.4959,0.4943,0.4951,0.4935,0.4953,0.4967,0.4981,0.4988,0.4995,0.498,0.4981,0.4974
0.5,0.42,0.42,0.52,0.5017,0.5044,0.4975,0.496,0.5006,0.4957,0.5029,0.5033,0.498,0.4976,0.4989,0.4959,0.4938,0.4942,0.4933,0.4937,0.4963,0.4968,0.499,0.4994,0.497,0.4981,0.4974
0.5,0.48,0.48,0.51,0.5,0.5056,0.4975,0.4967,0.5017,0.4948,0.5038,0.5033,0.4987,0.4973,0.4983,0.4951,0.4938,0.494,0.4929,0.4941,0.4954,0.4965,0.4983,0.4994,0.4973,0.4977,0.4971
0.5,0.58,0.58,0.52,0.505,0.5067,0.5058,0.4987,0.5033,0.499,0.505,0.5052,0.501,0.4997,0.4989,0.4974,0.495,0.4953,0.495,0.4931,0.4957,0.4972,0.499,0.4992,0.4977,0.498,0.4982
```
