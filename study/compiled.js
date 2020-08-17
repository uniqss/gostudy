/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.CommonStruct = (function() {

    /**
     * Properties of a CommonStruct.
     * @exports ICommonStruct
     * @interface ICommonStruct
     * @property {string|null} [userName] CommonStruct userName
     * @property {number|null} [userId] CommonStruct userId
     */

    /**
     * Constructs a new CommonStruct.
     * @exports CommonStruct
     * @classdesc Represents a CommonStruct.
     * @implements ICommonStruct
     * @constructor
     * @param {ICommonStruct=} [properties] Properties to set
     */
    function CommonStruct(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * CommonStruct userName.
     * @member {string} userName
     * @memberof CommonStruct
     * @instance
     */
    CommonStruct.prototype.userName = "";

    /**
     * CommonStruct userId.
     * @member {number} userId
     * @memberof CommonStruct
     * @instance
     */
    CommonStruct.prototype.userId = 0;

    /**
     * Creates a new CommonStruct instance using the specified properties.
     * @function create
     * @memberof CommonStruct
     * @static
     * @param {ICommonStruct=} [properties] Properties to set
     * @returns {CommonStruct} CommonStruct instance
     */
    CommonStruct.create = function create(properties) {
        return new CommonStruct(properties);
    };

    /**
     * Encodes the specified CommonStruct message. Does not implicitly {@link CommonStruct.verify|verify} messages.
     * @function encode
     * @memberof CommonStruct
     * @static
     * @param {ICommonStruct} message CommonStruct message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    CommonStruct.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.userName != null && Object.hasOwnProperty.call(message, "userName"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.userName);
        if (message.userId != null && Object.hasOwnProperty.call(message, "userId"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.userId);
        return writer;
    };

    /**
     * Encodes the specified CommonStruct message, length delimited. Does not implicitly {@link CommonStruct.verify|verify} messages.
     * @function encodeDelimited
     * @memberof CommonStruct
     * @static
     * @param {ICommonStruct} message CommonStruct message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    CommonStruct.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a CommonStruct message from the specified reader or buffer.
     * @function decode
     * @memberof CommonStruct
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {CommonStruct} CommonStruct
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    CommonStruct.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommonStruct();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.userName = reader.string();
                break;
            case 2:
                message.userId = reader.int32();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a CommonStruct message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof CommonStruct
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {CommonStruct} CommonStruct
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    CommonStruct.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a CommonStruct message.
     * @function verify
     * @memberof CommonStruct
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    CommonStruct.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.userName != null && message.hasOwnProperty("userName"))
            if (!$util.isString(message.userName))
                return "userName: string expected";
        if (message.userId != null && message.hasOwnProperty("userId"))
            if (!$util.isInteger(message.userId))
                return "userId: integer expected";
        return null;
    };

    /**
     * Creates a CommonStruct message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof CommonStruct
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {CommonStruct} CommonStruct
     */
    CommonStruct.fromObject = function fromObject(object) {
        if (object instanceof $root.CommonStruct)
            return object;
        var message = new $root.CommonStruct();
        if (object.userName != null)
            message.userName = String(object.userName);
        if (object.userId != null)
            message.userId = object.userId | 0;
        return message;
    };

    /**
     * Creates a plain object from a CommonStruct message. Also converts values to other types if specified.
     * @function toObject
     * @memberof CommonStruct
     * @static
     * @param {CommonStruct} message CommonStruct
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    CommonStruct.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults) {
            object.userName = "";
            object.userId = 0;
        }
        if (message.userName != null && message.hasOwnProperty("userName"))
            object.userName = message.userName;
        if (message.userId != null && message.hasOwnProperty("userId"))
            object.userId = message.userId;
        return object;
    };

    /**
     * Converts this CommonStruct to JSON.
     * @function toJSON
     * @memberof CommonStruct
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    CommonStruct.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return CommonStruct;
})();

$root.MsgWrapper = (function() {

    /**
     * Properties of a MsgWrapper.
     * @exports IMsgWrapper
     * @interface IMsgWrapper
     * @property {number|null} [MessageId] MsgWrapper MessageId
     * @property {Uint8Array|null} [Data] MsgWrapper Data
     */

    /**
     * Constructs a new MsgWrapper.
     * @exports MsgWrapper
     * @classdesc Represents a MsgWrapper.
     * @implements IMsgWrapper
     * @constructor
     * @param {IMsgWrapper=} [properties] Properties to set
     */
    function MsgWrapper(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * MsgWrapper MessageId.
     * @member {number} MessageId
     * @memberof MsgWrapper
     * @instance
     */
    MsgWrapper.prototype.MessageId = 0;

    /**
     * MsgWrapper Data.
     * @member {Uint8Array} Data
     * @memberof MsgWrapper
     * @instance
     */
    MsgWrapper.prototype.Data = $util.newBuffer([]);

    /**
     * Creates a new MsgWrapper instance using the specified properties.
     * @function create
     * @memberof MsgWrapper
     * @static
     * @param {IMsgWrapper=} [properties] Properties to set
     * @returns {MsgWrapper} MsgWrapper instance
     */
    MsgWrapper.create = function create(properties) {
        return new MsgWrapper(properties);
    };

    /**
     * Encodes the specified MsgWrapper message. Does not implicitly {@link MsgWrapper.verify|verify} messages.
     * @function encode
     * @memberof MsgWrapper
     * @static
     * @param {IMsgWrapper} message MsgWrapper message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    MsgWrapper.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.MessageId != null && Object.hasOwnProperty.call(message, "MessageId"))
            writer.uint32(/* id 1, wireType 0 =*/8).int32(message.MessageId);
        if (message.Data != null && Object.hasOwnProperty.call(message, "Data"))
            writer.uint32(/* id 2, wireType 2 =*/18).bytes(message.Data);
        return writer;
    };

    /**
     * Encodes the specified MsgWrapper message, length delimited. Does not implicitly {@link MsgWrapper.verify|verify} messages.
     * @function encodeDelimited
     * @memberof MsgWrapper
     * @static
     * @param {IMsgWrapper} message MsgWrapper message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    MsgWrapper.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a MsgWrapper message from the specified reader or buffer.
     * @function decode
     * @memberof MsgWrapper
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {MsgWrapper} MsgWrapper
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    MsgWrapper.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.MsgWrapper();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.MessageId = reader.int32();
                break;
            case 2:
                message.Data = reader.bytes();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a MsgWrapper message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof MsgWrapper
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {MsgWrapper} MsgWrapper
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    MsgWrapper.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a MsgWrapper message.
     * @function verify
     * @memberof MsgWrapper
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    MsgWrapper.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.MessageId != null && message.hasOwnProperty("MessageId"))
            if (!$util.isInteger(message.MessageId))
                return "MessageId: integer expected";
        if (message.Data != null && message.hasOwnProperty("Data"))
            if (!(message.Data && typeof message.Data.length === "number" || $util.isString(message.Data)))
                return "Data: buffer expected";
        return null;
    };

    /**
     * Creates a MsgWrapper message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof MsgWrapper
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {MsgWrapper} MsgWrapper
     */
    MsgWrapper.fromObject = function fromObject(object) {
        if (object instanceof $root.MsgWrapper)
            return object;
        var message = new $root.MsgWrapper();
        if (object.MessageId != null)
            message.MessageId = object.MessageId | 0;
        if (object.Data != null)
            if (typeof object.Data === "string")
                $util.base64.decode(object.Data, message.Data = $util.newBuffer($util.base64.length(object.Data)), 0);
            else if (object.Data.length)
                message.Data = object.Data;
        return message;
    };

    /**
     * Creates a plain object from a MsgWrapper message. Also converts values to other types if specified.
     * @function toObject
     * @memberof MsgWrapper
     * @static
     * @param {MsgWrapper} message MsgWrapper
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    MsgWrapper.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults) {
            object.MessageId = 0;
            if (options.bytes === String)
                object.Data = "";
            else {
                object.Data = [];
                if (options.bytes !== Array)
                    object.Data = $util.newBuffer(object.Data);
            }
        }
        if (message.MessageId != null && message.hasOwnProperty("MessageId"))
            object.MessageId = message.MessageId;
        if (message.Data != null && message.hasOwnProperty("Data"))
            object.Data = options.bytes === String ? $util.base64.encode(message.Data, 0, message.Data.length) : options.bytes === Array ? Array.prototype.slice.call(message.Data) : message.Data;
        return object;
    };

    /**
     * Converts this MsgWrapper to JSON.
     * @function toJSON
     * @memberof MsgWrapper
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    MsgWrapper.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return MsgWrapper;
})();

$root.UserLogin = (function() {

    /**
     * Properties of a UserLogin.
     * @exports IUserLogin
     * @interface IUserLogin
     * @property {string|null} [userName] UserLogin userName
     * @property {number|null} [userId] UserLogin userId
     * @property {ICommonStruct|null} [info] UserLogin info
     */

    /**
     * Constructs a new UserLogin.
     * @exports UserLogin
     * @classdesc Represents a UserLogin.
     * @implements IUserLogin
     * @constructor
     * @param {IUserLogin=} [properties] Properties to set
     */
    function UserLogin(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * UserLogin userName.
     * @member {string} userName
     * @memberof UserLogin
     * @instance
     */
    UserLogin.prototype.userName = "";

    /**
     * UserLogin userId.
     * @member {number} userId
     * @memberof UserLogin
     * @instance
     */
    UserLogin.prototype.userId = 0;

    /**
     * UserLogin info.
     * @member {ICommonStruct|null|undefined} info
     * @memberof UserLogin
     * @instance
     */
    UserLogin.prototype.info = null;

    /**
     * Creates a new UserLogin instance using the specified properties.
     * @function create
     * @memberof UserLogin
     * @static
     * @param {IUserLogin=} [properties] Properties to set
     * @returns {UserLogin} UserLogin instance
     */
    UserLogin.create = function create(properties) {
        return new UserLogin(properties);
    };

    /**
     * Encodes the specified UserLogin message. Does not implicitly {@link UserLogin.verify|verify} messages.
     * @function encode
     * @memberof UserLogin
     * @static
     * @param {IUserLogin} message UserLogin message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    UserLogin.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.userName != null && Object.hasOwnProperty.call(message, "userName"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.userName);
        if (message.userId != null && Object.hasOwnProperty.call(message, "userId"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.userId);
        if (message.info != null && Object.hasOwnProperty.call(message, "info"))
            $root.CommonStruct.encode(message.info, writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
        return writer;
    };

    /**
     * Encodes the specified UserLogin message, length delimited. Does not implicitly {@link UserLogin.verify|verify} messages.
     * @function encodeDelimited
     * @memberof UserLogin
     * @static
     * @param {IUserLogin} message UserLogin message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    UserLogin.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a UserLogin message from the specified reader or buffer.
     * @function decode
     * @memberof UserLogin
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {UserLogin} UserLogin
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    UserLogin.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.UserLogin();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.userName = reader.string();
                break;
            case 2:
                message.userId = reader.int32();
                break;
            case 3:
                message.info = $root.CommonStruct.decode(reader, reader.uint32());
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a UserLogin message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof UserLogin
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {UserLogin} UserLogin
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    UserLogin.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a UserLogin message.
     * @function verify
     * @memberof UserLogin
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    UserLogin.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.userName != null && message.hasOwnProperty("userName"))
            if (!$util.isString(message.userName))
                return "userName: string expected";
        if (message.userId != null && message.hasOwnProperty("userId"))
            if (!$util.isInteger(message.userId))
                return "userId: integer expected";
        if (message.info != null && message.hasOwnProperty("info")) {
            var error = $root.CommonStruct.verify(message.info);
            if (error)
                return "info." + error;
        }
        return null;
    };

    /**
     * Creates a UserLogin message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof UserLogin
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {UserLogin} UserLogin
     */
    UserLogin.fromObject = function fromObject(object) {
        if (object instanceof $root.UserLogin)
            return object;
        var message = new $root.UserLogin();
        if (object.userName != null)
            message.userName = String(object.userName);
        if (object.userId != null)
            message.userId = object.userId | 0;
        if (object.info != null) {
            if (typeof object.info !== "object")
                throw TypeError(".UserLogin.info: object expected");
            message.info = $root.CommonStruct.fromObject(object.info);
        }
        return message;
    };

    /**
     * Creates a plain object from a UserLogin message. Also converts values to other types if specified.
     * @function toObject
     * @memberof UserLogin
     * @static
     * @param {UserLogin} message UserLogin
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    UserLogin.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults) {
            object.userName = "";
            object.userId = 0;
            object.info = null;
        }
        if (message.userName != null && message.hasOwnProperty("userName"))
            object.userName = message.userName;
        if (message.userId != null && message.hasOwnProperty("userId"))
            object.userId = message.userId;
        if (message.info != null && message.hasOwnProperty("info"))
            object.info = $root.CommonStruct.toObject(message.info, options);
        return object;
    };

    /**
     * Converts this UserLogin to JSON.
     * @function toJSON
     * @memberof UserLogin
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    UserLogin.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return UserLogin;
})();

$root.LoginBase = (function() {

    /**
     * Properties of a LoginBase.
     * @exports ILoginBase
     * @interface ILoginBase
     * @property {string|null} [sessionId] LoginBase sessionId
     * @property {string|null} [userName] LoginBase userName
     * @property {string|null} [password] LoginBase password
     */

    /**
     * Constructs a new LoginBase.
     * @exports LoginBase
     * @classdesc Represents a LoginBase.
     * @implements ILoginBase
     * @constructor
     * @param {ILoginBase=} [properties] Properties to set
     */
    function LoginBase(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * LoginBase sessionId.
     * @member {string} sessionId
     * @memberof LoginBase
     * @instance
     */
    LoginBase.prototype.sessionId = "";

    /**
     * LoginBase userName.
     * @member {string} userName
     * @memberof LoginBase
     * @instance
     */
    LoginBase.prototype.userName = "";

    /**
     * LoginBase password.
     * @member {string} password
     * @memberof LoginBase
     * @instance
     */
    LoginBase.prototype.password = "";

    /**
     * Creates a new LoginBase instance using the specified properties.
     * @function create
     * @memberof LoginBase
     * @static
     * @param {ILoginBase=} [properties] Properties to set
     * @returns {LoginBase} LoginBase instance
     */
    LoginBase.create = function create(properties) {
        return new LoginBase(properties);
    };

    /**
     * Encodes the specified LoginBase message. Does not implicitly {@link LoginBase.verify|verify} messages.
     * @function encode
     * @memberof LoginBase
     * @static
     * @param {ILoginBase} message LoginBase message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    LoginBase.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.sessionId != null && Object.hasOwnProperty.call(message, "sessionId"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.sessionId);
        if (message.userName != null && Object.hasOwnProperty.call(message, "userName"))
            writer.uint32(/* id 2, wireType 2 =*/18).string(message.userName);
        if (message.password != null && Object.hasOwnProperty.call(message, "password"))
            writer.uint32(/* id 3, wireType 2 =*/26).string(message.password);
        return writer;
    };

    /**
     * Encodes the specified LoginBase message, length delimited. Does not implicitly {@link LoginBase.verify|verify} messages.
     * @function encodeDelimited
     * @memberof LoginBase
     * @static
     * @param {ILoginBase} message LoginBase message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    LoginBase.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a LoginBase message from the specified reader or buffer.
     * @function decode
     * @memberof LoginBase
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {LoginBase} LoginBase
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    LoginBase.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.LoginBase();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.sessionId = reader.string();
                break;
            case 2:
                message.userName = reader.string();
                break;
            case 3:
                message.password = reader.string();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a LoginBase message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof LoginBase
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {LoginBase} LoginBase
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    LoginBase.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a LoginBase message.
     * @function verify
     * @memberof LoginBase
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    LoginBase.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.sessionId != null && message.hasOwnProperty("sessionId"))
            if (!$util.isString(message.sessionId))
                return "sessionId: string expected";
        if (message.userName != null && message.hasOwnProperty("userName"))
            if (!$util.isString(message.userName))
                return "userName: string expected";
        if (message.password != null && message.hasOwnProperty("password"))
            if (!$util.isString(message.password))
                return "password: string expected";
        return null;
    };

    /**
     * Creates a LoginBase message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof LoginBase
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {LoginBase} LoginBase
     */
    LoginBase.fromObject = function fromObject(object) {
        if (object instanceof $root.LoginBase)
            return object;
        var message = new $root.LoginBase();
        if (object.sessionId != null)
            message.sessionId = String(object.sessionId);
        if (object.userName != null)
            message.userName = String(object.userName);
        if (object.password != null)
            message.password = String(object.password);
        return message;
    };

    /**
     * Creates a plain object from a LoginBase message. Also converts values to other types if specified.
     * @function toObject
     * @memberof LoginBase
     * @static
     * @param {LoginBase} message LoginBase
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    LoginBase.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults) {
            object.sessionId = "";
            object.userName = "";
            object.password = "";
        }
        if (message.sessionId != null && message.hasOwnProperty("sessionId"))
            object.sessionId = message.sessionId;
        if (message.userName != null && message.hasOwnProperty("userName"))
            object.userName = message.userName;
        if (message.password != null && message.hasOwnProperty("password"))
            object.password = message.password;
        return object;
    };

    /**
     * Converts this LoginBase to JSON.
     * @function toJSON
     * @memberof LoginBase
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    LoginBase.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return LoginBase;
})();

module.exports = $root;
