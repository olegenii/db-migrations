/**
 * @apiDefine ItemNotFoundError
 * @apiVersion 0.1.0
 * @apiError ItemNotFound The id of the item was not found.
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
 * @apiVersion 0.1.0
 * @apiError InternalError There is an internal error.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 505 Internal Server Error
 *     {
 *       "error":  InternalError"
 *     }
 */
// Internal Server Error

/**
 * @apiDefine admin User access only
 * @apiVersion 0.1.0
 * This optional description belong to to the group admin.
 */

/**
 * @api {get}  /items/:id Get item
 * @apiVersion 0.1.0
 * @apiName GetItem
 * @apiGroup Items
 *
 * @apiParam {Number} id Items unique ID.
 *
 * @apiSuccess {String} title  Item title.
 * @apiSuccess {Number} price  Item price.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "id": 1,
 *       "title": "Apple",
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
 * @apiPermission admin
 *
 * @apiParam {Number} id       Items unique ID.
 * @apiParam {String} [title]  Item title.
 * @apiParam {Number} [price]  Item price.
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
 * @api {get} /items Get all items
 * @apiVersion 0.1.0
 * @apiName GetAllItems
 * @apiGroup Items
 *
 * @apiSuccess {Object[]} items 		Items list.
 * @apiSuccess {Number}   items.id 		Items unique ID.
 * @apiSuccess {String}   items.title 	Item title.
 * @apiSuccess {Number}   items.price 	Item price.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     [{
 *	   	"id": 1,
 *		"title": "apple",
 *		"price": "10.5"
 *     }]
 *
 * @apiUse InternalError
 */
// Get all item information
