/*
 * @Description:
 * @Author: Y95201
 * @Date: 2022-12-19 10:09:37
 * @LastEditors: Y95201
 * @LastEditTime: 2022-12-20 10:49:21
 */
package good_service

import (
	"github.com/gin-gonic/gin"
	"myChat/models"
	"net/http"
	"strconv"
)

type sortSons struct {
	models.Goods

	TotalWeight  int
	TotalAmount  int
	TotalDeposit int
}

func UserOrderList(c *gin.Context) {
	//2:定金、3:发现好物 4:行情锁价 5定金订货

	UserId := c.PostForm("user_id")
	if len(UserId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数为空"})
		return
	}
	IntUserId, _ := strconv.Atoi(UserId)
	models.GetUserByFieldValue("id", IntUserId)
	SellGoods := models.GetGoodsBylist(IntUserId)

	//if len(SellGoods) > 0 {
	//}
	//fmt.Println(reflect.TypeOf(SellGoods))
	//var attrs = map[int]interface{}{}

	//for i, v := range SellGoods {
	//fmt.Println(i, "-", v)
	//childrenCount := 56
	//total_amount := 456
	//total_deposit := 789
	//attrs[i] = sortSons{Goods: v, TotalWeight: childrenCount, TotalAmount: total_amount, TotalDeposit: total_deposit}

	//}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  SellGoods,
	})

	//$userId = UsersModel::create()->where('id',$user_id)->val('id');
	//$sellGoods = GoodsModel::create()->where('goods.user_id',$userId)
	//->where('goods.state',1)
	//->where('goods.number','0','!=')
	//->where('goods.type',2)
	//->where('goods.created_at', date('Y-m-d H:i:s', strtotime('-1 day')), '>')
	//->field(' any_value(goods.id) as id,
	//any_value(goods.note) as contract,
	//	COUNT(goods.note) AS count,
	//	any_value(goods.created_at) as created_at,
	//	any_value(goods.type) as order_type')
	//->group('goods.note')->order('goods.created_at', 'DESC')->all(null);
	//if($sellGoods){
	//	foreach( $sellGoods as $key => $value ){
	//		$lists = $this->toAllArray(GoodsModel::create()->where('goods.note', $value->contract)
	//		->field('goods.id, goods.name, goods.texture, goods.spec, goods.price, goods.number, goods.weight_ton, goods.unit, goods.total_weight, goods.usage_time, goods.user_id , round(goods.price * goods.total_weight , 2) as order_money')->all(null));
	//
	//		$sellGoods[$key]->goods = $lists;
	//      $sellGoods[$key]->total_weight =  round(array_sum(array_column($lists, 'total_weight')),3);
	//		$sellGoods[$key]->total_amount =  round(array_sum(array_column($lists, 'order_money')),2);
	//		$sellGoods[$key]->total_deposit = $this->DepositAlgorithm($lists);
	//	}
	//}

	//$buyGoods = GoodsModel::create()->where('orders.user_id',$userId)
	//->where('goods.user_id','')
	//->where('goods.type',4)
	//->where('goods.state',1)
	//->where('goods.number','0','!=')
	//->where('goods.created_at', date('Y-m-d H:i:s', strtotime('-1 day')), '>')
	//->join('orders','orders.goods_id = goods.id','LEFT')
	//->field('any_value(goods.id) as id,
	//any_value(orders.contract) as contract,
	//COUNT(orders.contract) AS count,
	//any_value(goods.created_at) as created_at,
	//any_value(goods.type) as order_type')
	//->group('orders.contract')->order('goods.created_at', 'DESC')->all(null);
	//if( $buyGoods){
	//	foreach( $buyGoods as $key => $value ){
	//		$lists = $this->toAllArray(GoodsModel::create()->where('orders.contract', $value->contract)
	//		->join('orders','orders.goods_id = goods.id','LEFT')
	//		->field('goods.id, goods.name, goods.texture, goods.spec, goods.price, goods.number, goods.weight_ton, goods.unit, goods.total_weight, goods.usage_time, goods.user_id, round(goods.price * goods.total_weight , 2) as order_money')->all(null));
	//
	//		$buyGoods[$key]->goods = $lists;
	//		$buyGoods[$key]->total_weight = round(array_sum(array_column($lists, 'total_weight')),3);
	//		$buyGoods[$key]->total_amount = round(array_sum(array_column($lists, 'order_money')),2);
	//		$buyGoods[$key]->total_deposit = $this->DepositAlgorithm($lists);
	//	}
	//}

	//$goods['sellGoods'] = $sellGoods;
	//$goods['buyGoods'] = $buyGoods;

	//$queryBuild = new QueryBuilder();
	//$queryBuild->raw("SELECT DISTINCT `orders`.`contract`,`goods`.`created_at`,`goods`.`usage_time` FROM `goods` LEFT JOIN `orders` ON `goods`.`id` = `orders`.`goods_id` WHERE (`goods`.`state` = 1 AND `goods`.`number` != 0 AND `goods`.`company_id` = 0 AND `goods`.`type` = 5 AND `orders`.`user_id` = ?) AND `goods`.`deleted_at` IS NULL", [$userId]);
	//$good = DbManager::getInstance()->query($queryBuild, true, 'default')->getResult();
	//if($good){
	//	$contract = [];
	//	foreach ($good as $key => $value) {
	//		$time =  strtotime("+".$value['usage_time']." day", strtotime($value["created_at"]));
	//		if (time() < $time) {
	//			$contract[] = $value['contract'];
	//		}
	//	}
	//	$depositGoods = GoodsModel::create()->where('orders.contract', $contract, 'IN')
	//	->join('orders','orders.goods_id = goods.id','LEFT')
	//	->field('any_value(goods.id) as id,
	//	any_value(orders.contract) as contract,
	//	COUNT(orders.contract) AS count,
	//	any_value(goods.created_at) as created_at,
	//	any_value(goods.type) as order_type')
	//	->group('orders.contract')->order('goods.created_at', 'DESC')->all(null);
	//	$lists = [];
	//	foreach ($depositGoods as $key => $value){
	//		$lists = $this->toAllArray(GoodsModel::create()->where('orders.contract', $value->contract)
	//		->join('orders','orders.goods_id = goods.id','LEFT')
	//		->field('goods.id, goods.name, goods.texture, goods.spec, goods.price, goods.number, goods.weight_ton, goods.unit, goods.total_weight, goods.usage_time, goods.user_id, round(goods.price * goods.total_weight , 2) as order_money')->all(null));
	//
	//		$depositGoods[$key]->goods = $lists;
	//		$depositGoods[$key]->total_weight = round(array_sum(array_column($lists, 'total_weight')),3);
	//		$depositGoods[$key]->total_amount = round(array_sum(array_column($lists, 'order_money')),2);
	//		$depositGoods[$key]->total_deposit = $this->DepositAlgorithm($lists);
	//	}
	//	$goods['depositGoods'] = $depositGoods;
	//}else{
	//	$goods['depositGoods'] = [];
	//}
	return
}

//func ArraySumColumn(data interface{}, column string) (int)
//	// 定义一个 map 用于存储统计结果
//	m := make(map[string]int)
//
//	// 定义要统计的 key
//	key := column
//	// 使用嵌套的 for 循环遍历数组
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			for k, v := range data[i][j] {
//			// 如果是指定的 key，就将值相加
//			if k == key {
//				m[k] += v
//			}
//		}
//	}
//	return m[key]
//}

