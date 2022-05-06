/**
 * @apiDefine ItemNotFoundError
 *
 * @apiError ItemNotFound The id of the Item was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error":  ItemNotFound"
 *     }
 */
// Item not found error

/**
 * @apiDefine InternalError
 *
 * @apiError InternalError There is some internal error.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 505 Internal Server Error
 *     {
 *       "error":  InternalError"
 *     }
 */
// Internal Server Error

/**
 * @api {get}  /items/:id Get item info
 * @apiVersion 0.1.0
 * @apiName GetItem
 * @apiGroup Items
 *
 * @apiParam {Number} id Items unique ID.
 *
 * @apiSuccess {String} title  Title of the Item.
 * @apiSuccess {Number} price  Price of the Item.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "id": 1,
 *		 "title": "Apple",
 *       "price": "10.5"
 *     }
 *
 * @apiUse ItemNotFoundError
 * @apiUse InternalError
 */
// Get Item information

/**
 * @api {put}  /items Modify item
 * @apiVersion 0.1.0
 * @apiName PutItem
 * @apiGroup Items
 *
 * @apiParam {Number} id       Items unique ID.
 * @apiParam {String} [title]  Title of the Item.
 * @apiParam {Number} [price]  Price of the Item.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiUse InternalError
 */
// Modify Item information

/**
 * @api {delete}  /items/:id Delete item
 * @apiVersion 0.1.0
 * @apiName DeleteItem
 * @apiGroup Items
 *
 * @apiParam {Number} id       Items unique ID.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiUse InternalError
 * @apiUse ItemNotFoundError
 */
// Delete Item

/**
 * @api {get} /items Request all item information
 * @apiVersion 0.1.0
 * @apiName GetAllItem
 * @apiGroup Items
 *
 * @apiSuccess {Object[]} items 		List of the Item.
 * @apiSuccess {Number}   items.id 		Items unique ID.
 * @apiSuccess {String}   items.title 	Title of the Item.
 * @apiSuccess {Number}   items.price 	Price of the Item.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     [{
 *	   	"id": 1,
 *		"title": "apple", 
 *		"price": "10.5"
 *	   }]
 *
 * @apiUse InternalError
 */
// Get all item information
