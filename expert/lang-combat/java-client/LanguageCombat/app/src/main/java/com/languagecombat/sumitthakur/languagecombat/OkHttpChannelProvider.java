package com.languagecombat.sumitthakur.languagecombat;

/**
 * Created by sumitthakur on 21/12/18.
 */


import io.grpc.Internal;
import io.grpc.InternalServiceProviders;
import io.grpc.ManagedChannelProvider;
import io.grpc.internal.GrpcUtil;
import io.grpc.okhttp.OkHttpChannelBuilder;

/**
 * Provider for {@link OkHttpChannelBuilder} instances.
 */
@Internal
public final class OkHttpChannelProvider extends ManagedChannelProvider {

    @Override
    public boolean isAvailable() {
        return true;
    }

    @Override
    public int priority() {
        return (GrpcUtil.IS_RESTRICTED_APPENGINE
                || InternalServiceProviders.isAndroid(getClass().getClassLoader())) ? 8 : 3;
    }

    @Override
    public OkHttpChannelBuilder builderForAddress(String name, int port) {
        return OkHttpChannelBuilder.forAddress(name, port);
    }

    @Override
    public OkHttpChannelBuilder builderForTarget(String target) {
        return OkHttpChannelBuilder.forTarget(target);
    }
}
