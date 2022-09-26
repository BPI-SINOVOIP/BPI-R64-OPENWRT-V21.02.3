/* grilo-net-0.3.vapi generated by vapigen-0.34, do not modify. */

[CCode (cprefix = "GrlNet", gir_namespace = "GrlNet", gir_version = "0.3", lower_case_cprefix = "grl_net_")]
namespace GrlNet {
	[CCode (cheader_filename = "net/grl-net.h", type_id = "grl_net_wc_get_type ()")]
	public class Wc : GLib.Object {
		[CCode (has_construct_function = false)]
		public Wc ();
		public static GLib.Quark error_quark ();
		public void flush_delayed_requests ();
		public async bool request_async (string uri, GLib.Cancellable? cancellable, [CCode (array_length_cname = "length", array_length_pos = 2.1, array_length_type = "gsize")] out uint8[] content) throws GLib.Error;
		[CCode (finish_name = "grl_net_wc_request_finish")]
		[Version (since = "0.2.2")]
		public async bool request_with_headers_async (string uri, GLib.Cancellable? cancellable, ..., [CCode (array_length_cname = "length", array_length_pos = 2.1, array_length_type = "gsize")] out uint8[] content) throws GLib.Error;
		[CCode (finish_name = "grl_net_wc_request_finish")]
		[Version (since = "0.2.2")]
		public async bool request_with_headers_hash_async (string uri, GLib.HashTable<string,string>? headers, GLib.Cancellable? cancellable, [CCode (array_length_cname = "length", array_length_pos = 2.1, array_length_type = "gsize")] out uint8[] content) throws GLib.Error;
		[Version (since = "0.1.12")]
		public void set_cache (bool use_cache);
		[Version (since = "0.1.12")]
		public void set_cache_size (uint cache_size);
		public void set_log_level (uint log_level);
		public void set_throttling (uint throttling);
		[NoAccessorMethod]
		public bool cache { get; set construct; }
		[NoAccessorMethod]
		public uint cache_size { get; set construct; }
		[NoAccessorMethod]
		public uint loglevel { get; set; }
		[NoAccessorMethod]
		public uint throttling { get; set; }
		[NoAccessorMethod]
		public string user_agent { owned get; set construct; }
	}
	[CCode (cheader_filename = "net/grl-net.h", cprefix = "GRL_NET_WC_ERROR_", has_type_id = false)]
	public enum WcError {
		UNAVAILABLE,
		PROTOCOL_ERROR,
		AUTHENTICATION_REQUIRED,
		NOT_FOUND,
		CONFLICT,
		FORBIDDEN,
		NETWORK_ERROR,
		PROXY_ERROR,
		CANCELLED
	}
}