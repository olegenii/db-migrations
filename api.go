// Item not found error
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

// Get Item information
/**
 * @api {get}  Item/:id Request Item information
 * @apiName Ge Item
 * @apiGroup Item
 *
 * @apiParam {Number} id Items unique ID.
 *
 * @apiSuccess {String} title Firstname of the Item.
 * @apiSuccess {String} lastname  Lastname of the Item.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "firstname": "John",
 *       "lastname": "Doe"
 *     }
 *
 * @apiUse ItemNotFoundError
 */
// Modify Item information
/**
 * @api {put}  Item/ Modify Item information
 * @apiName Pu Item
 * @apiGroup Item
 *
 * @apiParam {Number} id          Items unique ID.
 * @apiParam {String} [firstname] Firstname of the Item.
 * @apiParam {String} [lastname]  Lastname of the Item.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *
 * @apiUse ItemNotFoundError
 */