# 分层有限状态机

## 什么是分层有限状态机？

分层有限状态机，简称“HFSM”，是为了解决“有限状态机”在面对多状态下复杂跳转带来的编程复杂性。采用分治法的方式，对状态进行分层，从而简化编程难度。


## “FSM”与“HFSM”实现多状态复杂情形下的对比

FSM实现效果如下：

![FSM](http://aisharing.com/wp/wp-content/uploads/2011/08/fsm_thumb.png)


HFSM实现效果如下：

![HFSM](http://aisharing.com/wp/wp-content/uploads/2011/08/hfsm_thumb.png)


从上面两个图的对比，可以发现，采用HFSM会大大简化编程的难度。每一层的内部状态转移与外部无关，完全解耦。层与层之间的状态转移也会同内部状态一样进行转移。