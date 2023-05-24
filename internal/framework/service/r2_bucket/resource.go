package r2_bucket

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &R2BucketResource{}
var _ resource.ResourceWithImportState = &R2BucketResource{}

func NewResource() resource.Resource {
	return &R2BucketResource{}
}

// R2BucketResource defines the resource implementation.
type R2BucketResource struct {
	client *cloudflare.API
}

func (r *R2BucketResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_r2_bucket"
}

func (r *R2BucketResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *R2BucketResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *R2BucketModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r2Bucket, err := r.client.CreateR2Bucket(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()),
		cloudflare.CreateR2BucketParameters{
			Name:         data.Name.ValueString(),
			LocationHint: data.Location.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Error creating R2 bucket", err.Error())
		return
	}
	data.ID = types.StringValue(r2Bucket.Name)
	data.Name = types.StringValue(r2Bucket.Name)
	data.Location = types.StringValue(r2Bucket.Location)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2BucketResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *R2BucketModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	r2Bucket, err := r.client.GetR2Bucket(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading R2 bucket", err.Error())
		return
	}
	data.ID = types.StringValue(r2Bucket.Name)
	data.Name = types.StringValue(r2Bucket.Name)
	data.Location = types.StringValue(r2Bucket.Location)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *R2BucketResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *R2BucketModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddError("Error updating R2 bucket", "Not implemented")
}

func (r *R2BucketResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *R2BucketModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteR2Bucket(ctx, cloudflare.AccountIdentifier(data.AccountID.ValueString()), data.ID.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Error deleting R2 bucket", err.Error())
		return
	}
}

func (r *R2BucketResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("Error importing R2 bucket", "Invalid ID specified. Please specify the ID as \"account_id/name\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}
